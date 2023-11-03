import { React, useState } from "react";

function Navbar(props) {
    const toggleHomepage = props.clickHandlers.toggleHomepage
    const toggleProfile = props.clickHandlers.toggleProfile;
    const toggleNewPost = props.clickHandlers.toggleNewPost;
    const toggleDMList = props.clickHandlers.toggleDMList;
    const toggleSearchUser = props.clickHandlers.toggleSearchUser;
    const toggleNotifications = props.clickHandlers.toggleNotifications;

return (
  <div className="flex flex-row justify-around bg-gradient-to-r from-indigo-300 to-purple-300 fixed w-screen">

    <div onClick={toggleHomepage} className="flex flex-row items-center hover:scale-90 transition-all ease-in-out align-middle cursor-pointer">
        <img
        src="./logo.png"
        className="w-12"
        />
        <h1 className="text-black hidden md:block text-lg">&nbsp; SpaceBook</h1>
    </div>


    <div className="flex flex-row md:items-center">
      <img
        src="./cp.png"
        onClick={toggleNewPost}
        className="w-12 cursor-pointer p-2 hover:p-3 transition-all ease-in-out"
      />
      <img
        src="./search.png"
        onClick={toggleSearchUser}
        className="w-12 cursor-pointer p-2 hover:p-3 transition-all ease-in-out"
      />
      <img
        src="./orbit.png"
        onClick={toggleNotifications}
        className="w-12 cursor-pointer p-2 hover:p-3 transition-all ease-in-out"
      />
      <img
        src="./wormhole.png"
        onClick={toggleDMList}
        className="w-12 cursor-pointer p-2 hover:p-3 transition-all ease-in-out"
      />
      <img
        src="./profile.png"
        onClick={toggleProfile}
        className="w-12 cursor-pointer p-2 hover:p-3 transition-all ease-in-out"
      />
    </div>
  </div>
);

      
      
      
      
      

}

export default Navbar;