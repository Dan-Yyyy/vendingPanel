import React from "react";
import styled from "styled-components";

const ButtonContent = styled.button`
  display: block;
  width: 100%;
  max-width: 400px;
  padding: 13px;
  margin: 0 auto 24px;
  border: 1px solid #A0694B;
  border-radius: 10px;
  color: white;
  background-color: ${props => props.color};
  font-size: 24px;
  font-family: "Roboto-Thin";
  cursor: pointer;

  ${ props => props.color === 'white' && `
    color: black;
  `}

  &:hover {
    color: black;
    background-color: white;
    ${ props => props.color === 'white' && `
      color: white;
      background-color: #A0694B;
    `}
  }
`;

export const Button = ({children, color, click}) => {
  return(
    <ButtonContent color={color} onClick={click}>
      {children}
    </ButtonContent>
  )
}