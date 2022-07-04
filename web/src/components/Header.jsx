import React from "react";
import { FiMenu } from 'react-icons/fi';
import styled from "styled-components";

const HeaderContainer = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 32px;
  
  .header_menu {
    color: #A0694B;
    stroke-width: 1;
    cursor: pointer;
  }
`;

const User = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 5px 10px;
  height: 40px;
  background-color: #CCB3A5;
  border-radius: 10px;
  cursor: pointer;

  img {
    width: 30px;
    height: 30px;
    border-radius: 30px;
  }

  p {
    margin-left: 6px;
    color: white;
    font-size: 18px;
    font-weight: 500;
  }
`;

export const Header = () => {

  return(
    <HeaderContainer>
        <FiMenu size={40} className="header_menu"/>
      <User>
        <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTQPrvDwVG49SBYvvDQI0IqEFnuPr-iMGT7UA&usqp=CAU" alt="userImage" />
        <p>User</p>
      </User>
    </HeaderContainer>
  )
}