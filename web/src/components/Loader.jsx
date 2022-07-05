import React from "react";
import styled, { keyframes } from "styled-components";

const rotate = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`;

const LoaderContainer = styled.div`
  position: absolute;
  z-index: 2;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.3);

  div {
    width: 75px;
    height: 75px;
    border-radius: 50%;
    border: 2px dashed #A0694B;
    animation: ${rotate} 2s linear infinite;
  }
`;

export const Loader = () => {
  return(
    <LoaderContainer>
      <div></div>
    </LoaderContainer>
  )
};