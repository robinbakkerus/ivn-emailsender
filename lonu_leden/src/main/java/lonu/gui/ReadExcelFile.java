package lonu.gui;

import java.io.File;
import java.time.LocalDate;
import java.util.List;
import java.util.Properties;
import java.util.Set;
import java.util.stream.Collectors;

import lonu.gui.data.State;
import lonu.model.User;
import lonu.util.ExcelUtils;
import lonu.util.FileUtils;
import lonu.util.SettingUtils;
import javafx.event.ActionEvent;
import javafx.geometry.Insets;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.scene.text.Text;
import javafx.stage.FileChooser;
import javafx.stage.Stage;

public class ReadExcelFile {

	private final Stage stage;
	private final State state = new State();
	private final FileUtils fileutils = new FileUtils();
	private final ExcelUtils excelUtils = new ExcelUtils();
	
	private Button generateButton;
	private Text resultText;

	public ReadExcelFile(Stage stage) {
		super();
		this.stage = stage;
	}

	public void exec() {
		stage.setTitle("Vergelijk LONU emails");
		this.readSetting();

		HBox hBox1 = this.buildSelectDumpFile("Select Lonu leden excel", this.state.getLonuEmails(), SettingUtils.DUMP_DIR);
		HBox hBox2 = this.buildSelectDumpFile("Select EmailPoet excel", this.state.getPoetEmails(), SettingUtils.SKIP_DIR);
		
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

	private HBox buildSelectDumpFile(String title, Set<User> targetSet, String propName) {
		FileChooser fileChooser = new FileChooser();
		
		fileChooser.setInitialDirectory(this.initialDir(this.settings(), propName));
		fileChooser.getExtensionFilters().addAll(new FileChooser.ExtensionFilter("Excel Files", "*.xlsx") );

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
	
	private void onFileSelected(FileChooser fileChooser, Text text, ActionEvent event, Set<User> targetSet, Properties settings, String propName) {
		File selectedFile = fileChooser.showOpenDialog(stage);
		settings.setProperty(propName, selectedFile.getParent());
		List<User> users = this.excelUtils.parseExcel(selectedFile.getAbsolutePath());
		text.setText(selectedFile.getAbsolutePath());
		targetSet.clear();
		targetSet.addAll(users);
		
		if (this.state.hasAllData()) {
			this.generateButton.setDisable(false);
		}
	}
	
	private void doGenerate() {
		Set<User> missingInLonu = this.state.getLonuEmails().stream().filter(e -> !this.state.getPoetEmails().contains(e)).collect(Collectors.toSet());
		showUsers(missingInLonu, "Emails die in Lonu lijst staan maar niet in EmailPoet lijst");
		
		Set<User> missingInPoet = this.state.getPoetEmails().stream().filter(e -> !this.state.getLonuEmails().contains(e)).collect(Collectors.toSet());
		showUsers(missingInPoet, "Emails die in EmailPoet lijst staan maar niet in Lonu leden lijst");

		Set<User> multEmailsPoet = this.state.getPoetEmails().stream().filter(e -> this.sameUserOtherEmail(e, this.state.getLonuEmails())).collect(Collectors.toSet());
		showUsers(multEmailsPoet, "Zelfde achternaam meerdere email adressen");

		Set<User> multEmailsLonu = this.state.getLonuEmails().stream().filter(e -> this.sameUserOtherEmail(e, this.state.getPoetEmails())).collect(Collectors.toSet());
		showUsers(multEmailsLonu, "Zelfde achternaam meerdere email adressen");
}

	private void showUsers(Set<User> missingInLonu, String header) {
		System.out.println("");
		System.out.println(header);
		for (User u : missingInLonu) {
			System.out.println(u);
		}
	}
	
	private boolean sameUserOtherEmail(User u, Set<User> otherSet) {
		return otherSet.stream().anyMatch(e -> e.getLastname().equals(u.getLastname()) && !e.getEmail().equals(u.getEmail()));
	}


}
