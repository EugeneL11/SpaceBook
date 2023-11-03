import { React, useState } from "react";
import MyProfile from "./MyProfile";
import Settings from "./Settings";
import FriendsList from "../Buddies/FriendsList";
import OtherProfile from "./OtherProfile";
function ProfileController(props) {
    const toggleLogin = props.toggleLogin
    const toggleHomepage = props.toggleHomepage
    const [profileState,setProfileState] = useState(<MyProfile toggleHomepage = {toggleHomepage} toggleFriendsList={toggleFriendsList} toggleSettings ={toggleSettings}/>)
    function toggleMyProfile(){
        setProfileState(<MyProfile toggleHomepage = {toggleHomepage} toggleFriendsList={toggleFriendsList} toggleSettings ={toggleSettings}/>)
    }
    function toggleFriendsList(){
        setProfileState(<FriendsList toggleOtherProfile = {toggleOtherProfile} toggleMyProfile ={toggleMyProfile} toggleFriendsList = {toggleFriendsList}/>)
    }
    function toggleSettings(){
        setProfileState(<Settings toggleMyProfile={toggleMyProfile} toggleLogin = {toggleLogin}/>);
    }
    function toggleOtherProfile(userID, backEvent){
        setProfileState(<OtherProfile userID ={userID} goBackScreen ={backEvent}/>);
    }
    return (<div>{profileState}</div>);

}

export default ProfileController;