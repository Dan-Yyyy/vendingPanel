import instanse from './instanse';

export const auth = {
  signIn(userData) {
    return instanse.post(`auth/sign-in`, userData)
    .then(response => {
      return {
        data: response.data,
        status: response.status,
        message: ''
      }
    })
  },
}