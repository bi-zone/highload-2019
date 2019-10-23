import scala.concurrent.ExecutionContextExecutor
import scala.util.Failure
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer
import akka.stream.scaladsl.{Flow, Framing, Sink, Tcp}
import akka.util.ByteString

object Solution {

  def main(args: Array[String]): Unit = {

    implicit val system: ActorSystem = ActorSystem()
    implicit val materializer: ActorMaterializer = ActorMaterializer()
    implicit val executionContext: ExecutionContextExecutor = system.dispatcher

    val host = "127.0.0.1"
    val port = 7777

    Tcp().bind(host, port)
      .runForeach { connection =>
        val address = s"${connection.remoteAddress.getAddress.getHostAddress}:${connection.remoteAddress.getPort}"
        println(s"New connection from $address")

        val printResult = Sink.foreach[String] { result =>
          println(s"Result for connection from $address\n$result\n")
        }

        val processConnectionFlow = Flow[ByteString]
          .via(Framing.delimiter(ByteString("\n"), maximumFrameLength = 1024, allowTruncation = true))
          .takeWhile(_.nonEmpty)

          // TODO: implement
          .map { e =>
            // println(s">> ${e.utf8String}")
            1
          }
          .reduce(_ + _)
          .map(total => s"total events received: $total")

          .alsoTo(printResult)
          .map(ByteString(_))

        connection.handleWith(processConnectionFlow)
      }
      .onComplete {
        case Failure(exception) =>
          System.err.println(s"Oops: ${exception.getMessage}")
          system.terminate()
        case _ =>
      }
  }
}
