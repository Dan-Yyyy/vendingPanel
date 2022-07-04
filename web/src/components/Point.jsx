import React from "react";
import { useDispatch } from "react-redux";
import styled from "styled-components";

import { visited } from './../redux/home/homeReduser';

const PointContainer = styled.div`
  display: flex;
  justify-content: space-between;
  padding: 16px 24px;
  margin-bottom: 16px;
  color: black;
  border-radius: 10px;
  border: 1px solid #A0694B;
`;

const PointAdress = styled.p`
  margin-bottom: 8px;
  font-family: ${props => props.family};
  font-size: 18px;

  &:last-child {
    margin-bottom: 0;
  }
`;

const RadioContainer = styled.div`
  display: flex;
  align-items: center;
`;

const Input = styled.input`
  vertical-align: top;
  margin: 0 3px 0 0;
  width: 17px;
  height: 17px;

  &:not(checked) { 
    position: absolute; 
    opacity: 0;
  }

  &:not(checked) + label {
    position: relative; 
    padding: 0 0 0 48px;
    display: block;
    height: 20px;
    border-radius: 13px;
  }

  &:not(checked) + label:before {
    content: '';
    position: absolute;
    top: -2px;
    left: 0;
    width: 48px;
    height: 24px;
    border-radius: 13px;
    border-color: #CCB3A5;
    background-color: #CCB3A5;
  }

  &:not(checked) + label:after {
    content: '';
    position: absolute;
    top: 0px;
    left: 2px;
    width: 20px;
    height: 20px;
    border-radius: 10px;
    background: #FFF;
    transition: all .2s;
  }

  &:checked + label:after { 
    left: 26px;
  }
`;

const Label = styled.label``;

export const Point = ({adress, mashine, isVisited, id}) => {
  const dispatch = useDispatch();

  const handleChange = () => {
    dispatch(visited(id));
  }
  return(
    <PointContainer style={isVisited ? {backgroundColor: '#A0694B', color: '#fff'} : null}>
      <div>
        <PointAdress family={'Roboto-Regular'}>{adress}</PointAdress>
        <PointAdress family={'Roboto-Medium'}>{mashine}</PointAdress>
      </div>
      <RadioContainer>
        <Input type="checkbox" checked={isVisited} onChange={handleChange} id={adress + mashine}/>
        <Label htmlFor={adress + mashine}></Label>
      </RadioContainer>
    </PointContainer>
  )
}