
# Shrinkr - Fast and Secure URL Shortener

Shrinkr is a high-performance URL shortener built with speed and security in mind. It leverages Redis and MongoDB as databases, along with a blazingly fast backend server written in Fiber. With Shrinkr, you can effortlessly create shortened URLs, set expiration times, and even protect your URLs with secret codes for added security.




## Tech Stack

**Client:** Cobra

**Server:** Fiber

Here's the [backend repo](https://github.com/varun7singh/Shrinkr-API)


## Features

* _Lightning-Fast Performance_: Shrinkr is built using Fiber, a high-performance web framework for Go.

* _Browser SSO Authentication_: Shrinkr makes user authentication simple and easy with browser Single Sign-On (SSO) using Google Authentication.

* _URL Expiry Management_: Set expiration times for your shortened URLs, ensuring they are accessible only for a specific duration.

* _Protected URLs_: Take control of your URLs by adding a secret code. Only users with the correct code can access the protected URLs.

* And many more to come... 
## Installation

Clone the project

```bash
    git clone https://github.com/varun7singh/Shrinkr
```

Navigate to the dir

```bash
    cd Shrinkr
```

Build the tool using Make

```bash
    make build
```

And Voila!! you're all set to use Shrinkr

```bash
    shrinkr login
```

  
    
## Feedback

If you have any feedback, please reach out to me varun7singh10@gmail.com

__Thanks!!__
