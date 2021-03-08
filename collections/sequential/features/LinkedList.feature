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

