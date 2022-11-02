package lonu;

import lonu.gui.ReadExcelFile;
import javafx.application.Application;
import javafx.stage.Stage;

public class LonuMain extends Application {

	@Override
	public void start(Stage stage) {
		ReadExcelFile selectFiles = new ReadExcelFile(stage);
		selectFiles.exec();
	}

	public static void main(String[] args) {
		launch();
	}
}
