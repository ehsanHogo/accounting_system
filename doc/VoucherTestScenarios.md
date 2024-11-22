# Accounting Voucher Scenarios

## Create:

### Scenario: User cannot leave the number field empty for an Accounting Voucher

- **Given** the user is on the "Create Accounting Voucher" page
- **When** the user leaves the "Number" field empty
  - **And** clicks the "Submit" button
- **Then** an error message "Voucher number cannot be empty" is displayed
  - **And** the Accounting Voucher is not created

### Scenario: User cannot enter a number longer than 64 characters for an Accounting Voucher

- **Given** the user is on the "Create Accounting Voucher" page
- **When** the user enters a number longer than 64 characters
  - **And** clicks the "Submit" button
- **Then** an error message "Accounting Voucher number cannot exceed 64 characters" is displayed
  - **And** the Accounting Voucher is not created

### Scenario: User cannot enter a duplicate number for an Accounting Voucher

- **Given** an Accounting Voucher with the number "123" already exists
- **When** the user enters the number "123"
  - **And** clicks the "Submit" button
- **Then** an error message "Accounting Voucher number must be unique" is displayed
  - **And** the Accounting Voucher is not created

### Scenario: DL is required and must be in database when the SL is DL-enabled

- **Given** the user is creating a Voucher Item
  - **And** its SL is DL-enabled
- **When** the user does not select a DL
  - **And** clicks the "Submit" button
- **Then** an error message "DL is mandatory for a SL that is DL-enabled" is displayed

  - **And** the Voucher Item is not created

- **When** the selected DL does not exist in the database
  - **And** the user clicks the "Submit" button
- **Then** an error message "The selected DL does not exist in the database" is displayed
  - **And** the Voucher Item is not created

### Scenario: DL must be empty if the SL is not DL-enabled

- **Given** the user is creating a Voucher Item whose SL is not DL-enabled
- **When** the user selects a DL
  - **And** clicks the "Submit" button
- **Then** an error message "DL must be empty for a voucher Item that its SL is not DL-enabled" is displayed
  - **And** the Voucher Item is not created

### Scenario: The SL must exist in the database and is mandatory

- **Given** the user is creating a Voucher Item
- **When** the user does not add SL
  - **And** clicks the "Submit" button
- **Then** an error message "SL is mandatory for voucher item" is displayed

  - **And** the Voucher Item is not created

- **When** the SL does not exist in the database
  - **And** the user clicks the "Submit" button
- **Then** an error message "The SL does not exist in the database" is displayed
  - **And** the Voucher Item is not created

### Scenario: User cannot enter invalid values for debit and credit

- **Given** the user is creating a Voucher Item
- **When** the user enters both debit and credit as zero or both as non-zero or one of them as negative
  - **And** clicks the "Submit" button
- **Then** an error message "One of the debit or credit must be greater than zero and the other must be zero" is displayed
  - **And** the Voucher Item is not created

### Scenario: Accounting Vouchers must be balanced

- **Given** the user is trying to create a voucher
  - **And** the total debits do not equal the total credits
- **When** the user clicks the "Submit" button
- **Then** an error message "The voucher must be balanced" is displayed
  - **And** the voucher is not submitted

### Scenario: An accounting voucher must have between 2 and 500 items

- **Given** the user is trying to create an accounting voucher
- **When** the user adds less than 2 items to the voucher
  - **And** clicks the "Submit" button
- **Then** an error message "A voucher must have at least 2 items" is displayed

  - **And** the voucher is not submitted

- **When** the user adds more than 500 items to the voucher
  - **And** clicks the "Submit" button
- **Then** an error message "A voucher cannot have more than 500 items" is displayed
  - **And** the voucher is not submitted

### Scenario: User can create new voucher entity

- **Given** the user wants to add a new voucher entity that is not present in the database
- **When** the user enters a valid and non-repeated number, also valid voucher item count
- **Then** a message "Voucher entity added successfully" is displayed
  - **And** the voucher entity is created

## Delete:

### Scenario: User cannot delete an accounting voucher if the requested version was different from the database

- **Given** the accounting voucher with id "0" and version "1" is in the database
- **When** the user tries to delete the accounting voucher with id "0" and version "0"
- **Then** an error message "This accounting voucher version is different from database" is displayed
  - **And** the entity is not deleted

### Scenario: User cannot delete a voucher entity if it is not in the database

- **Given** the voucher entity with id "123" is not in the database
- **When** the user tries to delete the entity with id "123"
- **Then** an error message "This voucher entity is not in database and cannot be deleted" is displayed

### Scenario: User can delete a voucher entity

- **Given** a voucher entity with id "0" and version "0" is in the database
- **When** the user wants to delete the voucher entity with id "0" and version "0"
- **Then** a message "Voucher entity version is deleted successfully from the database" is displayed
  - **And** the entity is deleted

### Scenario: Accounting Vouchers must be balanced

- **Given** the user is trying to delete items from a voucher
  - **And** the total debits do not equal the total credits
- **When** the user clicks the "Submit" button
- **Then** an error message "The voucher must be balanced" is displayed
  - **And** the voucher is not submitted

### Scenario: An accounting voucher must have at least 2 items

- **Given** the user is trying to delete items from an accounting voucher
- **When** the voucher has less than 2 items
  - **And** clicks the "Submit" button
- **Then** an error message "A voucher must have at least 2 items" is displayed
  - **And** the voucher is not submitted

## Insert:

(Empty for now)

## Update:

### Scenario: User cannot update an accounting voucher if the requested version was different from the database

- **Given** the accounting voucher with id "0" and version "1" is in the database
- **When** the user tries to update the accounting voucher with id "0" and version "0"
- **Then** an error message "This accounting voucher version is different from database" is displayed
  - **And** the entity is not updated

### Scenario: User cannot update a voucher entity if it is not in the database

- **Given** the voucher entity with id "123" is not in the database
- **When** the user tries to update the entity with id "123"
- **Then** an error message "This voucher entity is not in database and cannot be updated" is displayed

### Scenario: User can update a voucher entity

- **Given** a voucher entity with id "0" and version "0" is in the database
- **When** the user wants to update the voucher entity with id "0" and version "0"
- **Then** a message "Voucher entity version is updated successfully in the database" is displayed
  - **And** the entity is updated

### Scenario: Accounting Vouchers must be balanced

- **Given** the user is trying to update items from a voucher
  - **And** the total debits do not equal the total credits
- **When** the user clicks the "Submit" button
- **Then** an error message "The voucher must be balanced" is displayed
  - **And** the voucher is not submitted

### Scenario: An accounting voucher cannot have more than 500 items

- **Given** the user is trying to insert items from an accounting voucher
- **When** the voucher has more than 500 items
  - **And** clicks the "Submit" button
- **Then** an error message "A voucher cannot have more than 500 items" is displayed
  - **And** the voucher is not submitted

### Scenario: User cannot enter invalid values for debit and credit

- **Given** the user is creating a Voucher Item
- **When** the user enters both debit and credit as zero or both as non-zero or one of them as negative
  - **And** clicks the "Submit" button
- **Then** an error message "One of the debit or credit must be greater than zero and the other must be zero" is displayed
  - **And** the Voucher Item is not created

## Read:

### Scenario: User cannot read a voucher entity if it is not in the database

- **Given** the voucher entity with id "123" is not in the database
- **When** the user tries to read the entity with id "123"
- **Then** an error message "This voucher entity is not in the database and cannot be read" is displayed

### Scenario: User can read a voucher entity

- **Given** a voucher entity with id "0" is in the database
- **When** the user wants to read the voucher entity with id "0"
- **Then** a message "Voucher entity is read successfully from the database" is displayed
  - **And** the entity is passed to the user
