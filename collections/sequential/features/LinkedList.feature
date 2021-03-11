Feature: Singly linked list
  @SLLEmpty
  Scenario Outline: New list is empty
    Given linked list implementation is "<implementation>"
    When I create a new list
    Then head and tail are nil and size is zero
    Examples:
    |implementation  |
    |SinglyLinkedList|
    |DoublyLinkedList|
    |CircularlyLinkedList|

  Scenario Outline: Appending
    Given linked list implementation is "<implementation>"
    When I append items
      |value|
      |1    |
      |2    |
      |3    |
    Then Head is 1
    And Tail is 3
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

  Scenario Outline: Prepending
    Given linked list implementation is "<implementation>"
    When I prepend items
      |value|
      |1    |
      |2    |
      |3    |
    Then Head is 3
    And Tail is 1
    Examples:
      |implementation|
      |SinglyLinkedList|
      |DoublyLinkedList|
      |CircularlyLinkedList|

