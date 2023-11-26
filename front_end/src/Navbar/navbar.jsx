import { React, useState } from "react";

function NavbarIcon(props) {
    return (
        <div 
            onClick={props.click}
            className={`flex flex-col hover:-translate-y-1 transition-all ease-in-out align-middle justify-center cursor-pointer 
            text-black hover:after:text-xs hover:after:text-center ${props.label}`}
        >
            <img src={props.image} className="w-12 scale-50"/>
        </div>
    )
}

function Navbar(props) {
    const toggleHomepage = props.clickHandlers.toggleHomepage
    const toggleProfile = props.clickHandlers.toggleProfile;
    const toggleNewPost = props.clickHandlers.toggleNewPost;
    const toggleDMList = props.clickHandlers.toggleDMList;
    const toggleSearchUser = props.clickHandlers.toggleSearchUser;
    const toggleNotifications = props.clickHandlers.toggleNotifications;

    return (
    <div className="flex flex-row justify-around bg-gradient-to-r from-indigo-300 to-purple-300 fixed w-screen h-16 top-0 z-50">

        <div onClick={toggleHomepage} className="flex flex-row items-center hover:scale-90 transition-all ease-in-out align-middle cursor-pointer">
            <img src="./logo.png" className="w-12"/>
            <h1 className="text-black hidden md:block text-lg">&nbsp; SpaceBook</h1>
        </div>

        <div className="flex flex-row md:items-center">
            <NavbarIcon image="./cp.png" label="hover:after:content-['New_Post']" click={toggleNewPost}/>
            <NavbarIcon image="./search.png" label="hover:after:content-['Search']" click={toggleSearchUser}/>
            <NavbarIcon image="./orbit.png" label="hover:after:content-['Requests']" click={toggleNotifications}/>
            <NavbarIcon image="./wormhole.png" label="hover:after:content-['Messages']" click={toggleDMList}/>
            <NavbarIcon image="./profile.png" label="hover:after:content-['Profile']" click={toggleProfile}/>
        </div>

    </div>
    );
}

export default Navbar;