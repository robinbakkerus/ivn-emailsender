package lonu.util;

import java.io.File;
import java.io.FileInputStream;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.poi.ss.usermodel.Cell;
import org.apache.poi.ss.usermodel.CellType;
import org.apache.poi.ss.usermodel.Row;
import org.apache.poi.ss.usermodel.Sheet;
import org.apache.poi.ss.usermodel.Workbook;
import org.apache.poi.xssf.usermodel.XSSFWorkbook;

import lonu.model.User;

public class ExcelUtils {

	private int firstNameIndex;
	private int lastNameIndex;
	private int emailIndex;
	
	public List<User> parseExcel(String filename) {
		List<User> result = new ArrayList<>();

		FileInputStream file;
		try {
			file = new FileInputStream(new File(filename));
			Workbook workbook = new XSSFWorkbook(file);

			Sheet sheet = workbook.getSheetAt(0);
			this.getIndexes(sheet);

			for (int rownr = 2; rownr <= sheet.getLastRowNum(); rownr++) {
				Row row = sheet.getRow(rownr);
				String firstname = this.strValue(row, this.firstNameIndex);
				String lastname = this.strValue(row, this.lastNameIndex);
				String email = this.strValue(row, this.emailIndex);
				result.add(new User(firstname, lastname, email));
			}
			
			workbook.close();
		} catch (Exception e) {
			e.printStackTrace();
		}

		return result;
	}

	private void getIndexes(Sheet sheet) {
		Row row = sheet.getRow(1); //eerst kijken of deze geldt
		this.emailIndex = this.getColIndex(row, "email");
		if (this.emailIndex < 0) {
			row = sheet.getRow(0);
			this.emailIndex = this.getColIndex(row, "e-mail");
		} 
		
		this.firstNameIndex = this.getColIndex(row, "voornaam");
		this.lastNameIndex = this.getColIndex(row, "achternaam");
	}

	// ---------- private -------

	private int getColIndex(Row row, String header) {
		int index = 0;
		for (Cell cell : row) {
			if (CellType.STRING.equals(cell.getCellType())) {
				if (cell.getStringCellValue().toLowerCase().startsWith(header.toLowerCase())) {
					return index;
				}
			}
			index++;
		}
		return -1;
	}
	
	private String strValue(Row row, int index) {
		Cell cell = row.getCell(index);
		if (cell != null && CellType.STRING.equals(cell.getCellType())) {
			return cell.getStringCellValue();
		} else {
			return "";
		}
	}
}
