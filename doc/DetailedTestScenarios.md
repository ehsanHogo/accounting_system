# DL Entity Scenarios

## Create

### Scenario: User cannot leave the code field empty for a DL entity

- **Given** the user is on the "Create Detail" page
- **When** the user leaves the "Code" field empty
  - And clicks the "Submit" button
- **Then** an error message "DL entity code cannot be empty" is displayed
  - And the DL entity is not created

### Scenario: User cannot enter a code longer than 64 characters for a DL entity

- **Given** the user is on the "Create Detail" page
- **When** the user enters a code longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "DL entity code cannot exceed 64 characters" is displayed
  - And the DL entity is not created

### Scenario: User cannot enter a duplicate code for a DL entity

- **Given** a DL entity with the code "123" already exists
- **When** the user enters the code "123"
  - And clicks the "Submit" button
- **Then** an error message "DL entity code must be unique" is displayed
  - And the DL entity is not created

### Scenario: User cannot leave the title field empty for a DL entity

- **Given** the user is on the "Create Detail" page
- **When** the user leaves the "Title" field empty
  - And clicks the "Submit" button
- **Then** an error message "DL entity title cannot be empty" is displayed
  - And the DL entity is not created

### Scenario: User cannot enter a title longer than 64 characters for a DL entity

- **Given** the user is on the "Create Detail" page
- **When** the user enters a title longer than 64 characters
  - And clicks the "Submit" button
- **Then** an error message "DL entity title cannot exceed 64 characters" is displayed
  - And the DL entity is not created

### Scenario: User cannot enter a duplicate title for a DL entity

- **Given** a DL entity with the title "abc" already exists
- **When** the user enters the title "abc"
  - And clicks the "Submit" button
- **Then** an error message "DL entity title must be unique" is displayed
  - And the DL entity is not created

### Scenario: User can create a new DL entity

- **Given** the user wants to add a new DL entity that is not present in the database
- **When** the user enters a valid and non-repeated code and title
- **Then** a message "DL entity added successfully" is displayed
  - And the DL entity is created

## Delete

### Scenario: User cannot delete a DL entity if it is referenced

- **Given** the DL entity with id "123" is referenced in a voucher
- **When** the user tries to delete the entity with id "123"
- **Then** an error message "This DL entity is referenced and cannot be deleted" is displayed
  - And the entity is not deleted

### Scenario: User cannot delete a DL entity if the requested version was different from the database

- **Given** the DL entity with id "0" and version "1" is in the database
- **When** the user tries to delete the DL entity with id "0" and version "0"
- **Then** an error message "This DL entity version is different from the database" is displayed
  - And the entity is not deleted

### Scenario: User cannot delete a DL entity if it is not in the database

- **Given** the DL entity with id "123" is not in the database
- **When** the user tries to delete the entity with id "123"
- **Then** an error message "This DL entity is not in the database and cannot be deleted" is displayed

### Scenario: User can delete a DL entity

- **Given** a DL entity with id "0" and version "0" is in the database
- **When** the user wants to delete the DL entity with id "0" and version "0"
- **Then** a message "DL entity version is deleted successfully from the database" is displayed
  - And the entity is deleted

## Update

### Scenario: User cannot update a DL entity if the requested version was different from the database

- **Given** the DL entity with id "0" and version "1" is in the database
- **When** the user tries to update the DL entity with id "0" and version "0"
- **Then** an error message "This DL entity version is different from the database" is displayed
  - And the entity is not updated

### Scenario: User cannot update a DL entity if it is not in the database

- **Given** the DL entity with id "123" is not in the database
- **When** the user tries to update the entity with id "123"
- **Then** an error message "This DL entity is not in the database and cannot be updated" is displayed

### Scenario: User can update a DL entity

- **Given** a DL entity with id "0" and version "0" is in the database
- **When** the user wants to update the DL entity with id "0" and version "0"
- **Then** a message "DL entity version is updated successfully in the database" is displayed
  - And the entity is updated

## Read

### Scenario: User cannot read a DL entity if it is not in the database

- **Given** the DL entity with id "123" is not in the database
- **When** the user tries to read the entity with id "123"
- **Then** an error message "This DL entity is not in the database and cannot be read" is displayed

### Scenario: User can read a DL entity

- **Given** a DL entity with id "0" is in the database
- **When** the user wants to read the DL entity with id "0"
- **Then** a message "DL entity is read successfully from the database" is displayed
  - And the entity is passed to the user
