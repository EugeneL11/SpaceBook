class User {
    constructor(){
        this.userID = null;
        this.admin = false;
        this.pfp = null;
        this.userName = null;
        this.email = null;
        this.planet = null;
        this.bio = null;
        this.full_name = null;
    }
}

const currentUser = new User();
export default currentUser;