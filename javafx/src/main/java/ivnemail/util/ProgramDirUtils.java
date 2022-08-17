package ivnemail.util;

import java.io.File;
import java.net.URISyntaxException;

public class ProgramDirUtils {
	public static String getProgramDirectory() {
		if (runningFromJAR()) {
			return getCurrentJARDirectory();
		} else {
			return getCurrentProjectDirectory();
		}
	}

	//----- private ---
	
	private static String getJarName() {
		return new File(ProgramDirUtils.class.getProtectionDomain().getCodeSource().getLocation().getPath())
				.getName();
	}

	private static boolean runningFromJAR() {
		String jarName = getJarName();
		return jarName.contains(".jar");
	}

	private static String getCurrentProjectDirectory() {
		return new File("").getAbsolutePath();
	}

	private static String getCurrentJARDirectory() {
		try {
			return new File(ProgramDirUtils.class.getProtectionDomain().getCodeSource().getLocation().toURI()
					.getPath()).getParent();
		} catch (URISyntaxException exception) {
			exception.printStackTrace();
		}

		return null;
	}
}