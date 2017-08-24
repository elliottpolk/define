## define

Hi, I'm ***define***. I'm a simple CLI interface for performing a [Google](https://google.com) + **define** search. 

Just a quick note before you can use me, I do require a valid [CSE ID](https://cse.google.com/cse/all) and [API key](https://console.cloud.google.com/apis/api/customsearch.googleapis.com/overview).

### Installation

The creator did not include downloadable (it's a word, look it up 😁), so you'll need to install and build using [Go](https://golang.org/).

```
$ go get -u github.com/elliottpolk/define
```

### Usage

Since you'll need to configure me before use, the ```-c | -config``` flag should help. You should note that the sample *key* and *CSE ID* are not actually real. So... be sure to generate your own.

```
$ define -c
Please provide a valid Google API key:
If6VR9J8FNuozeCYe2PGjxb0qrgss4wk3PJlPbJ
Please provide a valid custom search engine ID ('cx') value:
256323651702721704984:UfRCV4tMrsL
```

To ***define*** a a phrase you'll need to use the ```-p | -phrase``` flag along with the phrase. If you're looking for more than one word, it should be wrapped in ```""```.

```
# single word phrase
$ define -p customer
define -p customer
INFO[0001] showing top 3 of 281000000 total results for terms define+customer

INFO[0001] Define customer: someone who buys goods or services from a business —
customer in a sentence.

INFO[0001] Definition of customer: General: A party that receives or consumes products (
goods or services) and has the ability to choose between different products and ...

INFO[0001] Customer definition, a person who purchases goods or services from another;
buyer; patron. See more.

# multi-word phrase
$ define -p "customer service"
INFO[0001] showing top 3 of 65600000 total results for terms define+customer service

INFO[0001] May 10, 2015 ... Customer service is the act of taking care of the customer's needs by providing
and delivering professional, helpful, high quality service and...

INFO[0001] The process of ensuring customer satisfaction with a product or service. Often,
customer service takes place while performing a transaction for the customer, ...

INFO[0001] Find the definition of customer service. What is good customer service? Excellent
customer service job resources.
```

I will also allow for some poor spelling. I'll be sure to give you the corrected ***define*** term and the resulting definitions.
 
```
$ define -p csutomer
INFO[0001] showing top 3 of 111000000 total results for corrected terms define+customer

INFO[0001] Define customer: someone who buys goods or services from a business —
customer in a sentence.

INFO[0001] Customer definition, a person who purchases goods or services from another;
buyer; patron. See more.

INFO[0001] Definition of customer value: The difference between what a customer gets from a
product, and what he or she has to give in order to get it.
```

I'm pretty dependant on Google at this point, so... if they get it wrong, I get it wrong. I also do **not** give the option to ***Search instead for***... yet.

## Thanks ~
