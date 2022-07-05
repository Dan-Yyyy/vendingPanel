import React from "react";
import styled from "styled-components";

const InputContent = styled.input`
  width: 100%;
  padding: 0 0 3px;
  border: none;
  border-bottom: 1px solid #626262;
  border-radius: 0;
  color: black;
  font-size: 18px;
  font-family: "Roboto-Regular";
  &:focus-visible {
    outline: none;
  }
  &:focus {
    outline: none;
  }
  &::placeholder {
    color: #626262;
    font-family: "Roboto-Thin";
  }
  &:focus {
    border-bottom: 2px solid black;
  }
`;

export const Input = (props) => {
  return(
    <InputContent {...props}/>
  )
}