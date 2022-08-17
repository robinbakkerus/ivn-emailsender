package ivnemail.gui;

import java.io.File;
import java.time.LocalDate;
import java.util.List;
import java.util.Properties;
import java.util.Set;
import java.util.stream.Collectors;

import ivnemail.gui.data.State;
import ivnemail.util.FileUtils;
import ivnemail.util.SettingUtils;
import javafx.event.ActionEvent;
import javafx.geometry.Insets;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.scene.text.Text;
import javafx.stage.FileChooser;
import javafx.stage.Stage;

public class ReadDumpFile {

	private final Stage stage;
	private final State state = new State();
	private final FileUtils fileutils = new FileUtils();
	
	private Button generateButton;
	private Text resultText;

	public ReadDumpFile(Stage stage) {
		super();
		this.stage = stage;
	}

	public void exec() {
		stage.setTitle("Generare IVN email list");
		this.readSetting();

		HBox hBox1 = this.buildSelectDumpFile("Select dump file", this.state.getEmails(), SettingUtils.DUMP_DIR);
		HBox hBox2 = this.buildSelectDumpFile("Select skip emails file", this.state.getSkipEmails(), SettingUtils.SKIP_DIR);
		
		this.generateButton = buildGenerateButton();
		this.resultText = new Text();
		
		VBox vbox = new VBox(10.0, hBox1, hBox2, this.generateButton, this.resultText);
		Scene scene = new Scene(vbox, 960, 600);
		stage.setScene(scene);
		stage.show();
	}

	//----------------
	
	private void readSetting() {
		this.fileutils.readSettings(this.state);
	}

	private Properties settings() {
		return this.state.getSettings();
	}
	
	private Button buildGenerateButton() {
		Button button = new Button("Generate");
		button.setMinWidth(200.0);
		button.setDisable(true);
		button.setOnAction(e -> this.doGenerate());
		return button;
	}

	private HBox buildSelectDumpFile(String title, Set<String> targetSet, String propName) {
		FileChooser fileChooser = new FileChooser();
		
		fileChooser.setInitialDirectory(this.initialDir(this.settings(), propName));
		fileChooser.getExtensionFilters().addAll(new FileChooser.ExtensionFilter("CSV Files", "*.csv") );

		Text text = new Text();
		text.setText(" ... ");
		
		Button button = new Button(title);
		button.setMinWidth(200.0);
		button.setOnAction(e -> this.onFileSelected(fileChooser, text, e, targetSet, this.settings(), propName));
		
		HBox hBox = new HBox(30.0, button, text);
		return hBox;
	}
	
	private File initialDir(Properties settings, String propName) {
		String value = settings.getProperty(propName);
		if (value != null) {
			return new File(value);
		} else {
			return new File("c:\\temp");
		}
	}
	
	private void onFileSelected(FileChooser fileChooser, Text text, ActionEvent event, Set<String> targetSet, Properties settings, String propName) {
		File selectedFile = fileChooser.showOpenDialog(stage);
		settings.setProperty(propName, selectedFile.getParent());
		List<String> content = this.fileutils.readFile(selectedFile.getAbsolutePath());
		text.setText(selectedFile.getAbsolutePath());
		targetSet.clear();
		targetSet.addAll(this.fileutils.getEmailAddresses(content));
		
		if (this.state.hasAllData()) {
			this.generateButton.setDisable(false);
		}
	}
	
	private void doGenerate() {
		Set<String> finalEmails = this.state.getEmails().stream().filter(e -> !this.state.getSkipEmails().contains(e)).collect(Collectors.toSet());
		
		String content = String.join(",\n", finalEmails);
		if (this.fileutils.saveFile(this.outputFilename(), content)) {
			this.state.saveSetting();
			String text = "Email adressen opgeslagen in " + this.outputFilename();
			this.resultText.setText(text);
		}
	}

	private String outputFilename() {
		LocalDate today = LocalDate.now();
		String dumpdir = this.settings().getProperty(SettingUtils.DUMP_DIR);
		String result = String.format("%s\\emails-%s-%s-%s.txt",dumpdir, today.getDayOfMonth(), today.getMonth(), today.getYear());
		return result;
	}
}
