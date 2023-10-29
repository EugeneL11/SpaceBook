import logo from './logo.svg';
import './App.css';
import FetchTest from './FetchTest'

import { useState, useTransition } from 'react';
import Login from './Login/login';
import Register from './Login/register';
import Homepage from './Homepage/homepage';
import NewPost from './Posts/NewPost';
import ProfileController from './Profile/ProfileController';
import SearchUsers from './Buddies/SearchUsers';
import Notifcations from './Buddies/Notifications';
import OtherProfile from './Profile/OtherProfile';
import ExpandedPost from './Posts/ExpandedPost';
import DMController from './Messages/DMController';
import Navbar from './Navbar/navbar';
function App() {
  const [navBar,setNavBar] = useState(null)
  const [screen, setScreen] = useState(<Login toggleHomepage = {showHomeScreen} toggleRegister = {showRegisterScreen}/>);
  const showNavBar = () =>{
    if(navBar === null){
      const clickHandlers = {
        toggleProfile : showMyProfile,
        toggleNewPost : showNewPost,
        toggleDMList : showDMList,
        toggleSearchUser : showSearchUser,
        toggleNotifications : showNotifications,
        toggleHomepage : showHomeScreen
      }
      setNavBar(<Navbar clickHandlers ={clickHandlers}/>)
    }
  }
  const hideNavBar = () =>{
    if(navBar !== null){
      setNavBar(null)
    }
  }
  function showOtherProfile(personID, backEvent){
    showNavBar();
    setScreen(<OtherProfile userID = {personID} goBackScreen = {backEvent}/>)
  }
  function showLoginScreen(){
    hideNavBar();
    setScreen(<Login toggleHomepage = {showHomeScreen} toggleRegister = {showRegisterScreen}/>)
  }
  function showRegisterScreen(){
    hideNavBar();
    setScreen(<Register toggleHomepage={showHomeScreen} toggleLogin = {showLoginScreen}/>)
  }
  
  function showHomeScreen(){
    showNavBar();
    setScreen(<Homepage toggleExpandPost = {expandPost}/>)
  }
  function showNewPost(){
    showNavBar();
    setScreen(<NewPost toggleHomepage={showHomeScreen}/>)
  }
  function expandPost(postID){
    showNavBar();
    setScreen(<ExpandedPost postID = {postID} toggleHomepage={showHomeScreen} toggleOtherProfile= {showOtherProfile} toggleExpandPost={expandPost}/>)
  }
  function showDMList(){
    showNavBar();
    setScreen(<DMController toggleHomePage = {showHomeScreen}/>)
  }
  function showMyProfile(){
    showNavBar();
    setScreen(<ProfileController toggleLogin={showLoginScreen} toggleHomepage = {showHomeScreen}/>)
  }
  function showSearchUser(){
    showNavBar();
    setScreen(<SearchUsers toggleSearchUser={showSearchUser} toggleOtherProfile={showOtherProfile} toggleHomepage = {showHomeScreen}/>)
  }
  function showNotifications(){
    showNavBar();
    setScreen(<Notifcations toggleNotifications ={showNotifications} toggleHomepage ={showHomeScreen} toggleOtherProfile ={showOtherProfile}/>)
  }
  return (
    <div>
      {navBar}
      {screen}
    </div>
  )
}

export default App;
