package ivnemail.util;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.List;
import java.util.Properties;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import ivnemail.gui.data.State;

public class FileUtils {

	public void readSettings(State state) {
		try {
			String runpath = ProgramDirUtils.getProgramDirectory();
			String settingsFname = runpath + "\\IvnEmailParser.properties";
			File settingsFile = new File(settingsFname);
			if (settingsFile.exists()) {
				Properties props = new Properties();
				InputStream is = new FileInputStream(settingsFile);
				props.load(is);
				state.setSettings(props);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
	
	public List<String> readFile(String fname) {
		try {
			Path path = Paths.get(fname);
			Stream<String> lines = Files.lines(path);
			List<String> result = lines.collect(Collectors.toList());
			lines.close();
			return result;
		} catch (IOException e) {
			e.printStackTrace();
			throw new RuntimeException(e.getMessage());
		}
	}

	/**
	 * here all duplicate email addresses are skipped, and also the first line if
	 * needed
	 * 
	 * @param inputLines
	 * @return
	 */
	public Set<String> getEmailAddresses(List<String> inputLines) {
		Set<String> result = new HashSet<>();

		List<String> lines = this.removeFirstLine(inputLines);
		int emailIndex = this.getEmailIndex(lines);

		for (String line : lines) {
			String items[] = this.getItems(line);
			result.add(items[emailIndex]);
		}

		return result;
	}

	public boolean saveFile(String fname, String content) {
		try (BufferedWriter writer = new BufferedWriter(new FileWriter(fname))) {
			writer.write(content);
			writer.close();
			return true;
		} catch (IOException e) {
			e.printStackTrace();
			return false;
		}
	}

	// ---------------------------------

	private List<String> removeFirstLine(List<String> lines) {
		String firstLine = lines.get(0).toLowerCase();
		if (firstLine.indexOf("systeem id") > -1) {
			lines.remove(0);
		}
		return lines;
	}

	private int getEmailIndex(List<String> lines) {
		String items[] = this.getItems(lines.get(0));
		for (int i = 0; i < items.length; i++) {
			if (items[i].indexOf("@") > 1) {
				return i;
			}
		}

		throw new RuntimeException("Kan geen kolom met email vinden");
	}

	private String[] getItems(String line) {
		line = line.replaceAll(",", ";");
		line += ";";
		return line.split(";");
	}
}
