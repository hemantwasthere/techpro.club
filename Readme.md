<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ClubTechPro/techpro.club">
    <img src="https://github.com/ClubTechPro/techpro.club/blob/master/assets/logos/logo-large.png" alt="Logo" >
  </a>

  <p align="center">
    We help students and freshers build their career portfolios by encouraging them to contribute to open-source projects
    <br />
    <br />
    <br />
    
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#founder">Founder</a></li>
    <li><a href="#core-team">Core Team</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

As a new programmer, you have a tremendous opportunity to learn and create fantastic softwares.

<a href="https://techpro.club">Techpro club</a>  is a platform to help you hone your programming skills and learn to collaborate with people early on in your career.

While you contribute to projects, ask questions, and share your knowledge, we track them down in one place and help you build a better portfolio.

Our core services to the open source projects will be free forever. Feel free to contribute to our repository

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

-   [Golang](https://go.dev/)
-   [Mongodb](https://www.mongodb.com/)
-   [Bootstrap](https://getbootstrap.com)
-   [AWS SES](https://aws.amazon.com/ses/)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

Get a copy of Golang and Mongodb on your system and follow the **Installation instructions**.

Optional - For emails, you will need an account with AWS described in **Prerequisites** or you may skip this step and comment out the sections in the code.

### Prerequisites

#### Github Project

Github project is needed to authenticate users. You may do so by visiting [https://github.com/settings/apps](https://github.com/settings/apps). Note the **Client ID** and **Client secret**. Also, set two **Callback URLs** for project and contributor authentication.

**Note** - Create a `Github App` and not `OAuth App` because the later doesn't permit multiple callback URLs.

#### AWS SES (Optional for emails)

You may create a free AWS account for a year with their terms and conditions.

For emails, we are using Amazon SES. Please go through the documentation https://docs.aws.amazon.com/ses/latest/dg/smtp-credentials.html.

Apart from that, generated SMTP Credentials will not work in place of `SES_ACCESS_ID` & `SES_ACCESS_SECRET`. You will need to go to **IAM > Users > Security Credentials > Create Access Key**. Note the credentials and replace the values in .env variables.

Also, remember to apply `AmazonSesSendingAccess` from **IAM > Users > Permissions > Add Permissions**

The above steps are before you unlock the sending limits in SES. Read https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html

### Installation

1. Clone the repo
    ```sh
    git clone https://github.com/ClubTechPro/techpro.club.git
    ```
2. Set Environment variables
   Rename the `src/.env.sample` to `src/.env` and replace with actual credentials obtained from the steps mentioned in Prerequisites.

3. Install Golang dependencies
    ```sh
    cd /path/to/project
    go mod tidy
    ```
4. Run the project
    ```sh
    go run main.go
    ```
    Open a browser and hit `http://localhost:8080`. If you want to host the application on your own server, <a href="hosting.md">read this</a>

<p align="right">(<a href="#top">back to top</a>)</p>

You will need to create OAuth projects in Github and Google and place these values in Callback URL

**Github Callback URL**
GB_CONTRIBUTOR_REDIRECT_URI=http://localhost:8080/contributors/github/callback
GB_PROJECT_REDIRECT_URI=http://localhost:8080/projects/github/callback

**Google Callback URL**
GO_CONTRIBUTOR_REDIRECT_URI=http://localhost:8080/contributors/google/callback

Edit
However, we donot use GB_PROJECT_REDIRECT_URI anymore. You can safely use the other two

<!-- USAGE EXAMPLES -->

<!-- ROADMAP -->

## Roadmap

-   [x] **v 0.1**
    -   [x] User Authentication
    -   [x] Contributor preferences
    -   [x] Projects Create, Read, Update
    -   [x] Registration welcome email
    -   [x] Master templates
-   [x] **v 0.2**
    -   [x] Profile management - avatar, delete account, etc.
    -   [x] Project feeds
    -   [x] Notification for contributors
    -   [x] Notification for projects
    -   [x] Fetch Github profile
    -   [x] Bookmarks and Reactions
-   [ ] **v 0.3**
    -   [ ] Project invite members
    -   [ ] Project manage members
    -   [ ] Share on social media
    -   [ ] Feed filters
-   [ ] **v 0.4**
    -   [ ] Project analytics
-   [ ] **v 0.5**
    -   [ ] Badges
-   [ ] **v 0.6**
    -   [ ] Integration with Gitlab
-   [ ] **v 0.7**
    -   [ ] Integration with Bitbucket
-   [ ] **v 0.8**
    -   [ ] Integration with Azure

More to follow. Perhaps, once these are done, we will move it to some other section for better tracking.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Techproclub -

-   Twitter: [@ClubTechpro](https://twitter.com/ClubTechpro)
-   Email: hello@techpro.club

Links -

-   Project Link: [https://github.com/ClubTechPro/techpro.club](https://github.com/ClubTechPro/techpro.club)
-   Application Link : [Techpro.club](https://techpro.club)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

-   [Best Readme Template](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#top">back to top</a>)</p>

## Founders
-   [Chilarai](https://github.com/chilarai)
-   [Rohit Verma](https://github.com/RohitV5)
