package lonu.model;

import java.util.Objects;

public class User {

	private String firstname;
	private String lastname;
	private String email;
	
	
	public User() {
		super();
	}
	
	public User(String firstname, String lastname, String email) {
		super();
		this.firstname = firstname.toLowerCase();
		this.lastname = lastname.toLowerCase();;
		this.email = email.toLowerCase();;
	}

	public String getFirstname() {
		return firstname;
	}
	public void setFirstname(String firstname) {
		this.firstname = firstname.toLowerCase();;
	}
	public String getLastname() {
		return lastname;
	}
	public void setLastname(String lastname) {
		this.lastname = lastname.toLowerCase();;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email.toLowerCase();;
	}

	@Override
	public int hashCode() {
		return Objects.hash(email);
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		User other = (User) obj;
//		return Objects.equals(email, other.email) && Objects.equals(firstname, other.firstname)
//				&& Objects.equals(lastname, other.lastname);
		return Objects.equals(email, other.email);
	}

	@Override
	public String toString() {
		return email + "  (" + firstname + ", " + lastname + ")" ;
	}
	
	
}
