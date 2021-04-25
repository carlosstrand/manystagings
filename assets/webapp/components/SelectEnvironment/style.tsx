import styled from 'styled-components';
import ListComp from '@material-ui/core/List';
import ListItemComp from '@material-ui/core/ListItem';
import ListItemTextComp from '@material-ui/core/ListItemText';
import ListItemAvatarComp from '@material-ui/core/ListItemAvatar';
import AvatarComp from '@material-ui/core/Avatar';
import ListItemIconComp from '@material-ui/core/ListItemIcon';
import AddIconComp from '@material-ui/icons/Add';

import LayersIconComp from '@material-ui/icons/Layers';

export const List = styled(ListComp).attrs({
  dense: false,
})`
  width: 100%;
`;

export const ListItem = styled(ListItemComp)``;

export const ListItemText = styled(ListItemTextComp)``;

export const ListItemAvatar = styled(ListItemAvatarComp)``;

export const Avatar = styled(AvatarComp).attrs({
  variant: 'square',
})`
  border-radius: 8px;
  background-color: #f6f6f6;
`;

export const EnvIcon = styled(LayersIconComp)`
  color: #0C2461;
`;

export const ListItemIcon = styled(ListItemIconComp)``;

// Add List Item

export const AddListItemText = styled(ListItemTextComp)`
  color: #9d9d9d;
  transition: color 0.5s;
`;

export const AddAvatar = styled(Avatar).attrs({
  variant: 'square',
})`
  background-color: #f6f6f6;
  color: #9d9d9d;
  transition: all 0.5s;
`;

export const AddIcon = styled(AddIconComp)``;

export const AddListItem = styled(ListItemComp)`
  &:hover {
    ${AddListItemText} {
      color: #14131a;
    }
    ${AddAvatar} {
      background-color: #0C2461;
      color: #fff;
    }
  }
`;
