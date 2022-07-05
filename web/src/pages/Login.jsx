import React, { useState } from "react";
import styled from "styled-components";
import { GrClose } from 'react-icons/gr';
import { useSelector } from "react-redux";

import backgroundImage from './../assets/images/background.png';
import { Button } from "../components/Button";
import { FormLogIn } from '../components/FormLogIn';
import { TitleBlock } from "../components/TitleBlock";
import { Error } from "../components/Error";
import { Loader } from './../components/Loader';

const LoginContainer = styled.div`
  position: relative;
  display: flex;
  width: 100%;
  height: 100vh;
  overflow: hidden;
`;

const ButtonContainer = styled.div`
  position: absolute;
  left: 0;
  width: 100%;
  height: 100%;
  padding: 0 24px;
  background-image: url(${backgroundImage});
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
  transition: all .2s linear;

  ${props => props.formShow === true && `
    left: -100%;
  `}
`;

const FormContainer = styled.div`
  position: absolute;
  left: 100%;
  width: 100%;
  padding: 24px;
  margin-top: 20vh;
  transition: all .2s linear;

  ${props => props.formShow === true && `
    left: 0;
  `}
`;

export const Login = () => {
  const error = useSelector(state => state.auth);
  const [show, setShow] = useState(false)

  const showForm = () => {
    setShow(!show);
  }

  return(
    <LoginContainer>
      <ButtonContainer formShow={show}>
        <div style={{paddingTop: '65vh'}}>
          <Button color={'white'} click={showForm}>Войти</Button>
        </div>
      </ButtonContainer>   
      <FormContainer formShow={show}>
        <TitleBlock title="ВХОД" Icon={ GrClose } click={showForm} size={ 22 }/>
        {
          error.status === "fulfilled" && error.isAuth === false
          ? <Error position='static'>{error.message}</Error>
          : null
        }
        <FormLogIn/>
      </FormContainer>
        {
          error.status === "pending"
          ? <Loader/>
          : null
        }
    </LoginContainer>
  )
}