package lonu.gui.data;

import java.util.HashSet;
import java.util.Properties;
import java.util.Set;

import lonu.model.User;
import lonu.util.SettingUtils;

public class State {
	
	private Set<User> lonuEmails = new HashSet<>();
	private Set<User> poetEmails = new HashSet<>();
	private SettingUtils settingUtils;

	public State() {
		this.settingUtils = new SettingUtils();
	}

	public boolean hasAllData() {
		return !this.lonuEmails.isEmpty() && !this.poetEmails.isEmpty();
	}

	public Set<User> getLonuEmails() {
		return lonuEmails;
	}

	public void setLonuEmails(Set<User> emails) {
		this.lonuEmails = emails;
	}

	public Set<User> getPoetEmails() {
		return poetEmails;
	}

	public void setPoetEmails(Set<User> skipEmails) {
		this.poetEmails = skipEmails;
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
