import { useState, useTransition } from "react";
import Login from "./Login/login";
import Register from "./Login/register";
import Homepage from "./Homepage/homepage";
import NewPost from "./Posts/NewPost";
import ProfileController from "./Profile/ProfileController";
import SearchUsers from "./Buddies/SearchUsers";
import Notifications from "./Buddies/Notifications";
import OtherProfile from "./Profile/OtherProfile";
import ExpandedPost from "./Posts/ExpandedPost";
import DMController from "./Messages/DMController";
import Navbar from "./Navbar/navbar";
import Background from "./Background/background";
import FriendsList from "./Buddies/FriendsList";
import ImageDemo from "./ImageDemo";

function App() {
    const [navBar, setNavBar] = useState(false);
    const [screen, setScreen] = useState(<Login toggleHomepage={showHomeScreen} toggleRegister={showRegisterScreen} />);
    const clickHandlers = {
        toggleProfile: showMyProfile,
        toggleNewPost: showNewPost,
        toggleDMList: showDMList,
        toggleSearchUser: showSearchUser,
        toggleNotifications: showNotifications,
        toggleHomepage: showHomeScreen,
    };
    const showNavBar = () => {
        setNavBar(true);
    };
    const hideNavBar = () => {
        setNavBar(false);
    };
    function showOtherProfile(personID, backEvent) {
        showNavBar();
        setScreen(<OtherProfile userID={personID} goBackScreen={backEvent} goDMList={showDMList} />);
    }
    function showLoginScreen() {
        hideNavBar();
        setScreen(<Login toggleHomepage={showHomeScreen} toggleRegister={showRegisterScreen} />);
    }
    function showRegisterScreen() {
        hideNavBar();
        setScreen(<Register toggleHomepage={showHomeScreen} toggleLogin={showLoginScreen} />);
    }
    function showHomeScreen() {
        showNavBar();
        setScreen(
            <Homepage
                toggleExpandPost={expandPost}
                toggleOtherProfile={showOtherProfile}
                toggleHomePage={showHomeScreen}
            />
        );
    }
    function showNewPost() {
        showNavBar();
        setScreen(<NewPost toggleHomepage={showHomeScreen} />);
    }
    function expandPost(postID) {
        hideNavBar();
        setScreen(
            <ExpandedPost
                postID={postID}
                toggleHomepage={showHomeScreen}
                toggleOtherProfile={showOtherProfile}
                toggleExpandPost={expandPost}
            />
        );
    }
    function showDMList() {
        showNavBar();
        setScreen(null);
        setTimeout(() => {
            setScreen(<DMController toggleHomePage={showHomeScreen} />);
        }, 0);
    }
    function showMyProfile() {
        showNavBar();
        setScreen(null);
        // set timeout of 1 ms
        setTimeout(() => {
            setScreen(<ProfileController toggleLogin={showLoginScreen} toggleHomepage={showHomeScreen} />);
        }, 0);
    }
    function showSearchUser() {
        showNavBar();
        setScreen(
            <SearchUsers
                toggleSearchUser={showSearchUser}
                toggleOtherProfile={showOtherProfile}
                toggleHomepage={showHomeScreen}
            />
        );
    }
    function showFriendsList() {
        // victor fix
        hideNavBar();
        setScreen(<FriendsList toggleOtherProfile={showOtherProfile} toggleHomepage={showHomeScreen} />);
    }
    function showNotifications() {
        showNavBar();
        setScreen(
            <Notifications
                toggleNotifications={showNotifications}
                toggleHomepage={showHomeScreen}
                toggleOtherProfile={showOtherProfile}
            />
        );
    }
    return (
        <div>
            <Background className="!-z-20" />
            {navBar ? <Navbar clickHandlers={clickHandlers} /> : null}
            <div className={navBar ? "mt-20" : "mt=0"}>{screen}</div>
        </div>
        //  <ImageDemo/>
    );
}

export default App;
