# SL Entity Scenarios

## Create

### Scenario: User cannot leave the code field empty for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user leaves the "Code" field empty
  - And clicks the "Submit" button
- **Then** an error message "SL entity code cannot be empty" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a code longer than 64 characters for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user enters a code longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "SL entity code cannot exceed 64 characters" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a duplicate code for an SL entity

- **Given** an SL entity with the code "789" already exists
- **When** the user enters the code "789"
  - And clicks the "Submit" button
- **Then** an error message "SL entity code must be unique" is displayed
  - And the SL entity is not created

### Scenario: User cannot leave the title field empty for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user leaves the "Title" field empty
  - And clicks the "Submit" button
- **Then** an error message "SL entity title cannot be empty" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a title longer than 64 characters for an SL entity

- **Given** the user is on the "Create Detail" page
- **When** the user enters a title longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "SL entity title cannot exceed 64 characters" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a duplicate title for an SL entity

- **Given** an SL entity with the title "123" already exists
- **When** the user enters the title "123"
  - And clicks the "Submit" button
- **Then** an error message "SL entity title must be unique" is displayed
  - And the SL entity is not created

### Scenario: User can create a new SL entity

- **Given** the user wants to add a new SL entity that is not present in the database
- **When** the user enters a valid and non-repeated code and title
- **Then** a message "SL entity added successfully" is displayed
  - And the SL entity is created

## Delete

### Scenario: User cannot delete an SL entity if it is referenced

- **Given** the SL entity with id "789" is referenced in a voucher
- **When** the user tries to delete the entity with id "789"
- **Then** an error message "This entity is referenced and cannot be deleted" is displayed
  - And the entity is not deleted

### Scenario: User cannot delete an SL entity if the requested version was different from the database

- **Given** the SL entity with id "0" and version "1" is in the database
- **When** the user tries to delete the SL entity with id "0" and version "0"
- **Then** an error message "This SL entity version is different from the database" is displayed
  - And the entity is not deleted

### Scenario: User cannot delete an SL entity if it is not in the database

- **Given** the SL entity with id "123" is not in the database
- **When** the user tries to delete the entity with id "123"
- **Then** an error message "This SL entity is not in the database and cannot be deleted" is displayed

### Scenario: User can delete an SL entity

- **Given** an SL entity with id "0" and version "0" is in the database
- **When** the user wants to delete the SL entity with id "0" and version "0"
- **Then** a message "SL entity version is deleted successfully from the database" is displayed
  - And the entity is deleted

## Update

### Scenario: User cannot update an SL entity if it is referenced

- **Given** the SL entity with id "789" is referenced in a voucher
- **When** the user tries to update the entity with id "789"
- **Then** an error message "This entity is referenced and cannot be updated" is displayed
  - And the entity is not updated

### Scenario: User cannot update an SL entity if the requested version was different from the database

- **Given** the SL entity with id "0" and version "1" is in the database
- **When** the user tries to update the SL entity with id "0" and version "0"
- **Then** an error message "This SL entity version is different from the database" is displayed
  - And the entity is not updated

### Scenario: User cannot update an SL entity if it is not in the database

- **Given** the SL entity with id "123" is not in the database
- **When** the user tries to update the entity with id "123"
- **Then** an error message "This SL entity is not in the database and cannot be updated" is displayed

### Scenario: User can update an SL entity

- **Given** an SL entity with id "0" and version "0" is in the database
- **When** the user wants to update the SL entity with id "0" and version "0"
- **Then** a message "SL entity version is updated successfully in the database" is displayed
  - And the entity is updated

### Scenario: User cannot leave the code field empty for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user update the "Code" to empty field
  - And clicks the "Submit" button
- **Then** an error message "SL entity code cannot be empty" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a code longer than 64 characters for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user update a code to longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "SL entity code cannot exceed 64 characters" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a duplicate code for an SL entity

- **Given** an SL entity with the code "789" already exists
- **When** the user update the code to "789"
  - And clicks the "Submit" button
- **Then** an error message "SL entity code must be unique" is displayed
  - And the SL entity is not created

### Scenario: User cannot leave the title field empty for an SL entity

- **Given** the user is on the "Create SL" page
- **When** the user update the "Title" to a empty field
  - And clicks the "Submit" button
- **Then** an error message "SL entity title cannot be empty" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a title longer than 64 characters for an SL entity

- **Given** the user is on the "Create Detail" page
- **When** the user update a title to longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "SL entity title cannot exceed 64 characters" is displayed
  - And the SL entity is not created

### Scenario: User cannot enter a duplicate title for an SL entity

- **Given** an SL entity with the title "123" already exists
- **When** the user update the title to "123"
  - And clicks the "Submit" button
- **Then** an error message "SL entity title must be unique" is displayed
  - And the SL entity is not created

## Read

### Scenario: User cannot read an SL entity if it is not in the database

- **Given** the SL entity with id "123" is not in the database
- **When** the user tries to read the entity with id "123"
- **Then** an error message "This SL entity is not in the database and cannot be read" is displayed

### Scenario: User can read an SL entity

- **Given** an SL entity with id "0" is in the database
- **When** the user wants to read the SL entity with id "0"
- **Then** a message "SL entity is read successfully from the database" is displayed
  - And the entity is passed to the user
