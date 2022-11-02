package lonu.util;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.Properties;

public class SettingUtils {

	public static final String SETTINGS_NAME = "lonuParser.properties";
	public static final String DUMP_DIR = "dump.dir";
	public static final String SKIP_DIR = "skip.dir";
	
	private String runpath;
	private FileUtils fileUtils = new FileUtils();
	private Properties properties = new Properties();
	
	
	public SettingUtils() {
		this.readSettings();
	}

	public String getRunpath() {
		return runpath;
	}

	public Properties getProperties() {
		return properties;
	}

	public void saveSetting() {
		try (OutputStream output = new FileOutputStream(this.settingsFname())) {
			this.properties.store(output, null);
			System.out.println(properties);

		} catch (IOException io) {
			io.printStackTrace();
		}
	}
	
	public String getSetting(String propName) {
		return this.properties.getProperty(propName);
	}
	
	public void setRunpath(String runpath) {
		this.runpath = runpath;
	}

	public void setProperties(Properties properties) {
		this.properties = properties;
	}

	private void readSettings() {
		try {
			this.runpath = ProgramDirUtils.getProgramDirectory();
			String settingsFname = this.settingsFname();
			File settingsFile = new File(settingsFname);
			if (settingsFile.exists()) {
				InputStream is = new FileInputStream(settingsFile);
				this.properties.load(is);
			} else {
				this.fileUtils.saveFile(settingsFname, "");
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	private String settingsFname() {
		String settingsFname = this.runpath + "\\" + SETTINGS_NAME;
		return settingsFname;
	}
}
