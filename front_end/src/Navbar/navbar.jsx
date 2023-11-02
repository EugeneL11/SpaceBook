import { React, useState } from "react";

function Navbar(props) {
    const toggleHomepage = props.clickHandlers.toggleHomepage
    const toggleProfile = props.clickHandlers.toggleProfile;
    const toggleNewPost = props.clickHandlers.toggleNewPost;
    const toggleDMList = props.clickHandlers.toggleDMList;
    const toggleSearchUser = props.clickHandlers.toggleSearchUser;
    const toggleNotifications = props.clickHandlers.toggleNotifications;
    return (<div className="flex flex-row justify-between">
        <div onClick={toggleHomepage}>Logo(go homepage)</div>
        <div onClick={toggleNewPost}>New Post</div>
        <div onClick={toggleSearchUser}> Search users</div>
        <div onClick={toggleNotifications}> Notifications</div>
        <div onClick={toggleDMList}>Direct Messages</div>
        <div onClick={toggleProfile}>My Profile</div>
    </div>);

}

export default Navbar;
