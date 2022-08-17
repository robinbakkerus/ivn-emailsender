package ivnemail.gui.data;

import java.util.HashSet;
import java.util.Properties;
import java.util.Set;

import ivnemail.util.SettingUtils;

public class State {
	
	private Set<String> emails = new HashSet<>();
	private Set<String> skipEmails = new HashSet<>();
	private SettingUtils settingUtils;

	public State() {
		this.settingUtils = new SettingUtils();
	}

	public boolean hasAllData() {
		return !this.emails.isEmpty() && !this.skipEmails.isEmpty();
	}

	public Set<String> getEmails() {
		return emails;
	}

	public void setEmails(Set<String> emails) {
		this.emails = emails;
	}

	public Set<String> getSkipEmails() {
		return skipEmails;
	}

	public void setSkipEmails(Set<String> skipEmails) {
		this.skipEmails = skipEmails;
	}

	public String getRunpath() {
		return this.settingUtils.getRunpath();
	}

	public void saveSetting() {
		this.settingUtils.saveSetting();
	}
	
	public Properties getSettings() {
		return this.settingUtils.getProperties();
	}
	
	public void setSettings(Properties properties) {
		this.settingUtils.setProperties(properties);
	}
	
	// ------- private ----


}
