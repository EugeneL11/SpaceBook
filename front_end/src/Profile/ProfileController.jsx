import { React, useState } from "react";
import MyProfile from "./MyProfile";
import Settings from "./Settings";
import FriendsList from "../Buddies/FriendsList";
import OtherProfile from "./OtherProfile";

// the design patterns for all the profiles
function ProfileController(props) {
    const toggleLogin = props.toggleLogin
    const toggleHomepage = props.toggleHomepage
    const togglePost = props.togglePost
    const [profileState,setProfileState] = useState(<MyProfile togglePost={togglePost} toggleHomepage = {toggleHomepage} toggleFriendsList={toggleFriendsList} toggleSettings ={toggleSettings}/>)
    function toggleMyProfile(){
        setProfileState(<MyProfile togglePost={togglePost} toggleHomepage = {toggleHomepage} toggleFriendsList={toggleFriendsList} toggleSettings ={toggleSettings}/>)
    }
    function toggleFriendsList(){
        setProfileState(<FriendsList toggleOtherProfile = {toggleOtherProfile} toggleMyProfile ={toggleMyProfile} toggleFriendsList = {toggleFriendsList}/>)
    }
    function toggleSettings(){
        setProfileState(<Settings toggleMyProfile={toggleMyProfile} toggleLogin = {toggleLogin}/>);
    }
    function toggleOtherProfile(userID, backEvent){
        setProfileState(<OtherProfile togglePost={togglePost} userID ={userID} goBackScreen ={backEvent}/>);
    }
    return (<div>{profileState}</div>);
}

export default ProfileController;