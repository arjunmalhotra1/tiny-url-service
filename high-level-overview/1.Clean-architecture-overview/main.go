package main

/*
What does clean architecture means and how can we build our software that satisfies clean architecture.
The idea of clean architecture has been taken from book clean architecture by Bob Martin.
Clean Architecture extends some of the ideas to a software design mindset.

// Se 1.png
How can we design software that is decoupled, very easy to test, and minimize our
dependence on external things like a database or UI or any external framework that our core business logic doesn't depend on?
1.png is the clean architecture diagram.

What kind of software does the clean architecture produce?
1. Software that is independent from Frameworks.
2. Software that has isolated (and testable) Business logic.
3. Software that is independent of UI & Database. UI & Database live in the outermost part.
4. Software that is independent of any external agencies.
Like we are not reliant on any sought of vendor, or anything to build the core business logic that the application is
trying to satisfy.

Why is it important?
1. Code is decoupled.
2. Code is very easy to test.
3. We have really good separation of concerns.
4. Code is easy to change because it's highly modular.
5. Makes our live way way way easier.

What not to do.
People write code that is very jammed together. Are jamming database calls, to business logic inside of a handler that is
accepting calls via some http.We have all the stuff just muddled together now and hardest thing to do now is,
if we want to change the business logic, we have no way to test it without running my entire http server.
To see if that change propagated through properly. We can follow some patterns.

Some of the major patterns that clean Architecture advocates are:
1. Dependency Inversion/ Dependency Rule -
If you see 1.png you will notice that the arrows flow inwards on the circle. What this means is that nothing on the
inner circle, can reference anything from the outer circle. Dependencies flow inwards.
We want to keep our entities like our objects, aggregates, our functions don't rely on anything like a database.
Things are around the business logic, are all kind of at the core. Strictly input and output kind of stuff like,
methods, functions, business logic of our application is at the very center and doesn't depend on anything.
It doesn't care what database we use. Database doesn't care how data is represented. Database doesn't care if we represent it with a
data with a webapp or a mobile app.

Next layer above Entities is Application Business Rules.
The further out we go in the layers the more concrete are the implementation details.
Towards the outer layer we have the practical database implementation.
We have out HTTP handlers, or UI or external pieces that calls out to the code in the center (see 1.png)
and then just returns an answer from the center.
The way we do this is, we have separation of layers with interfaces.
Easiest one to think about is, say we have some object, and we have some business logic with this object.
And say we want to persist it somewhere. Worst thing we want to do is, jam in direct
database sql code, all bottled together with an entity or inside an entity, or method of our entity.
Do not do that. We cannot test it. It is a bad design. It will make us strictly coupled with whatever database we pick.
If we have some separation with the interfaces we can use repository (Domain Driven Design) which is like an interface
to a datastore. We can have an interface that says "getById(id)" and it takes some id and returns a payload or an
entity.
Our entities don't have to care how they are stored at all because this is done at the outer layer.
So we can have a service layer that uses a repository interface to get and persist these kind of core domain objects.
And these core domain objects can have their business functionality and we can have this simple mapping layer between
how we want to store the data and how we want the data to be represented in the core domain.
Biggest take way from clean architecture is,
"Dependencies flow inwards, things from the outer edges can import and rely on things in the inner circle
but nothing in the center can rely anything out of it."
So we would never want the core of our domain to rely and depend on some concrete detail like which database we choose.
Imagine if want to switch from MySQL, to Mongo or from Postgres DB to couchDB, if our practical database implementation
is all wrapped ou together, with our business logic, inside our handler, we might have to rewrite our entire application,
v/s what we are trying to do is make our code modular and following this pattern, is we really want to make the
surface area as small as possible if we want to make a change and as long as we satisfy the existing interfaces nothing our code needs to change. We merely put a new implementation behind an existing interface.
2. Interface Adapters.
3. Attention to detail.
*/
func main() {

}
