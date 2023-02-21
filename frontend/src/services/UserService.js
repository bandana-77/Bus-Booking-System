import axios from 'axios';

const BUS_API_BASE_URL = "http://localhost:9080/";

class UserService {

    getBuses(){
        return axios.get(BUS_API_BASE_URL);
    }

    // getUsers(){
    //     return axios.get(USER_API_BASE_URL);
    // }

    // createUser(user){
    //     return axios.post(USER_API_BASE_URL, user);
    // }

    // getUserById(userId){
    //     return axios.get(USER_API_BASE_URL + '/' + userId);
    // }

    // updateUser(user, userId){
    //     return axios.put(USER_API_BASE_URL + '/' + userId, user);
    // }

    // deleteUser(userId){
    //     return axios.delete(USER_API_BASE_URL + '/' + userId);
    // }
}

export default new UserService()