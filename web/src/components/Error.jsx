import React from "react";
import styled from "styled-components";
import { BiError } from 'react-icons/bi';
 
const ErrorContent = styled.div`
  display: flex;
  align-items: center;
  position: ${props => props.position};
  top: 27px;
  color: red;
  font-size: 13px;

  ${props => props.position === 'static' && `
    margin-bottom: 8px;
  `}

  .icon {
    margin-right: 3px;
  }
`;

export const Error = ({children, position = 'absolute'}) => {
  return(
    <ErrorContent position={position}>
      <BiError size={12} className="icon"/>
      {children}
    </ErrorContent>
  )
}