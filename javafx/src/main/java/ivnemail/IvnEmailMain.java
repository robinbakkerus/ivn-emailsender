package ivnemail;

import ivnemail.gui.ReadDumpFile;
import javafx.application.Application;
import javafx.stage.Stage;

public class IvnEmailMain extends Application {

	@Override
	public void start(Stage stage) {
		ReadDumpFile selectFiles = new ReadDumpFile(stage);
		selectFiles.exec();
	}

	public static void main(String[] args) {
		launch();
	}
}
