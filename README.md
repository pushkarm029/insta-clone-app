![Insta-Clone-App](https://drive.google.com/uc?export=view&id=1z6BjGrL8-iiryYssnh1qSzKV4ODgVg7T)


## **🔥 Introduction**

**Insta-Clone-App** is a full-stack social media clone for learning and demonstrating product flows similar to Instagram. It includes authentication, profile pages, post upload, likes, comments, explore/search surfaces, reels, and a Go API backed by Firebase services.

Frontend is deployed using Github Pages ->  
<a href="https://pushkarm029.github.io/insta-clone-app" target="_blank">pushkarm029.github.io/insta-clone-app</a>

## **Project Status**

This repository is in active maintenance. The current maintenance focus is keeping the React frontend, Go backend, Firebase integration notes, local development workflow, and test coverage usable for portfolio review and future feature work.

The original build happened in 2023, with a 2026 refresh focused on documentation, local verification, dependency hygiene, and security cleanup.


Full Stack Application Will be ***Deployed soon***.

## **💥 Features**
- Create an account and sign in with Firebase Authentication.
- Upload image posts with captions.
- Browse posts from followed users on the home feed.
- Discover public posts through the explore page.
- Search users and open profile pages from search results.
- Like and comment on posts.
- View reels and a chill-zone section for experimental mini-games.
- Run a Go backend API with Firebase Firestore and Storage integrations.

## **🛠️ Local Development** :

### Frontend

1. Open your terminal and then type
    ```shell
    git clone https://github.com/Pushkarm029/insta-clone-app
    ```
2. cd into the folder
    ```shell
    cd insta-clone-app/
    ```
3. cd into the frontend folder
   ```shell
   cd frontend
   ```
4. install the required dependencies for frontend
    ```shell
    npm install
    ```
5. To start the application on localhost
    ```shell
    npm start
    ```
6. To deploy the frontend application on Github Pages
    ```shell
    npm run deploy
    ```
7. To run frontend tests locally
    ```shell
    CI=true npm test -- --watchAll=false
    ```
8. cd into backend folder
   ```shell
   cd ../backend
   ```
9. install the required dependencies for backend
   ```shell
   go get -u ./...
   ```
10. Start the server on :8080
    ```shell
    go run main.go
    ```

**Note : You need to restart backend server after every change in any .go file.**


## **❤️ Learnings** :

- I have learned how to use components, props, and state to create interactive user interfaces. 
- I have also gained a better understanding of how to structure a React project, including the use of ES6 syntax and styling with CSS.
- I have learned how to implement React Redux to make a global state and then use it anywhere.
- I have also become familiar with the React Router library for creating dynamic routes in my web applications. 
- I have also gained experience in deploying React applications to production.
- I have also gained a better understanding of Firebase Authentication, Storage and Firestore.
- Got Better Understanding of GO while implementing this as backend in this app.
- I learned how to iterate over data present in Firestore, converging it into JSON, and making HTTP requests using the Go-Gin Framework.
- Created REST API's using GO to perform CRU (CREATE, READ and UPDATE) operation.
- I learned how to build Unity games for the web and then implement them in a full-stack application.
## **⛑️ Maintenance** :

Feel free to open issue to add a feature request or report any BUG. It will be appreciated from the depth of my heart❤️.

## **📅 Future**

- Add infinite scroll and Lazy Loading.
- Implement the delete operation in the REST API.
- Lots of Bugs and css needs to be fixed.
- Jest Framework will be used for testing.
- Firebase will be used for Real Time Messaging Options.
- More Creativity will be added from my side.
- Make it more responsive.
- Migrate to TypeScript.
- Implement VideoCall Feature using WebRTC.
- Docker will be used for containerization.
