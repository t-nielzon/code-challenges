package main

/*
Address Book by State

Given a string with friends to visit in different states:
"John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Sal Carpenter, 73 6th Street, Boston MA"

Produce a result that sorts the names by state and lists the name of the state
followed by the name of each person residing in that state (people's names sorted).

When the result is printed:
Massachusetts
.....^John Daggett 341 King Road Plymouth Massachusetts
.....^Sal Carpenter 73 6th Street Boston Massachusetts
^Virginia
.....^Alice Ford 22 East Broadway Richmond Virginia

Where ^ represents a space.

The resulting string will be:
"Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia"

Notes:
- There can be a blank last line in the given string of addresses.
- The tests only contain CA, MA, OK, PA, VA, AZ, ID, IN for states.

States:
- AZ: Arizona
- CA: California
- ID: Idaho
- IN: Indiana
- MA: Massachusetts
- OK: Oklahoma
- PA: Pennsylvania
- VA: Virginia
*/

func AddressBookByState(str string) string {
}