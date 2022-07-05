import styled from 'styled-components';

const TitleContainer = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 32px;
  
  .icon {
    color: ${props => props.color};
    cursor: pointer;
    stroke-width: 1;
  }
`;

const Title = styled.h1`
  color: #000;
  font-size: 32px;
  font-weight: 400;
  text-transform: uppercase;
`;

export const TitleBlock = ({title, Icon, click = null, size = 36, color = "#A0694B"}) => {
  return(
    <TitleContainer onClick={click}>
      <Title>{title}</Title>
      <Icon size={size} className="icon" color={color}/>
    </TitleContainer>
  )
}
