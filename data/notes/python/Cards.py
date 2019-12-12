#!/usr/bin/python
# Think Python, Ch. 18 Inheritance

import random

class Card:
    """Represents a poker card"""

    # Integers to names mappings (Class variables)
    suit_names = [ 'Clubs', 'Diamonds', 'Hearts', 'Spades' ]
    rank_names = [  None, 'Ace',
                    '2', '3', '4', '5', '6', '7', '8', '9', '10',
                    'Jack', 'Queen', 'King' ]

    def __init__( self, suit = 0, rank = 2 ):
        # (Instance variables)
        self.suit = suit
        self.rank = rank

    def __str__( self ):
        return "%s of %s" % \
        (
            # Use the attribute rank from the object self as an index
            # into the list rank_names from the class Card.
            Card.rank_names[self.rank],
            Card.suit_names[self.suit]
        )

    def __lt__( self, other ):
        """Overloads less than (<)"""

        # check the ranks
        if self.rank < other.rank: return True
        if self.rank > other.rank: return False

        # ranks are the same ... check suits
        return self.suit < other.suit

class Deck:
    """Represents a deck of cards"""

    def __init__( self ):
        """Initializes the Deck with 52 cards"""
        self.cards = []
        for suit in range(4):
            for rank in range(1, 14):
                card = Card(suit, rank) # Deck HAS-A Card
                self.cards.append(card)

    def __str__( self ):
        """Returns a string representation of the deck"""
        out = []
        for card in self.cards:
            out.append(str(card))
        return " | ".join(out)

    # A veneer [dyha] method
    def shuffle( self ):
        """Shuffles the cards in this deck"""
        random.shuffle(self.cards)

# Inherit methods from Deck class
class Hand(Deck): # Hand IS-A kind of Deck
    """Represents a hand of cards"""

    # override method from Deck
    def __init__ ( self, label = '' ):
        self.cards = []
        self.label = label

if __name__ == '__main__':

    queen_of_diamonds = Card( 1, 12 )
    jack_of_diamonds = Card( 1, 11 )
    if jack_of_diamonds < queen_of_diamonds:
        print "%s is lower than %s" % (jack_of_diamonds, queen_of_diamonds)

    deck = Deck()
    deck.shuffle()
    print
    print "Shuffled deck: "
    print deck

    hand = Hand('new hand')
    print
    print hand.label
