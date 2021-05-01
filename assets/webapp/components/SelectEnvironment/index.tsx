import React, { useState } from 'react';
import Environment from '../../types/environment';
import { Link } from 'react-router-dom';
import {
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Avatar,
  AddIcon,
  // Add List Item
  AddListItem,
  AddAvatar,
  AddListItemText,
  EnvIcon,
} from './style';
import CreateEditEnvironmentModal from '../CreateEditEnvironmentModal';


interface SelectEnvironmentProps {
  environments: Environment[];
  refetch: () => void;
}

interface SelectEnvironmentListItemProps {
  env: Environment;
}

const SelectEnvironmentListItem = ({ env }: SelectEnvironmentListItemProps) => {
  const EnvLink = (props: any) => <Link to={`/environments/${env.id}/applications`} {...props} />;
  return (
    <ListItem key={env.id} button component={EnvLink}>
      <ListItemAvatar>
        <Avatar>
          <EnvIcon />
        </Avatar>
      </ListItemAvatar>
      <ListItemText>{env.name}</ListItemText>
    </ListItem>
  );
};

const SelectEnvironment = (props: SelectEnvironmentProps) => {
  const { refetch } = props;
  const [showCreateEnvModal, setShowCreateModal] = useState(false);
  return (
    <>
      <CreateEditEnvironmentModal
        id={null}
        open={showCreateEnvModal}
        onClose={() => setShowCreateModal(false)}
        refetch={refetch}
      />
      <List>
        {props.environments.map((env) => <SelectEnvironmentListItem key={env.id} env={env} />)}
        <AddListItem button onClick={() => setShowCreateModal(true)}>
          <ListItemAvatar>
            <AddAvatar>
              <AddIcon />
            </AddAvatar>
          </ListItemAvatar>
          <AddListItemText>Create Environment</AddListItemText>
        </AddListItem>
      </List>
    </>
  );
};

export default SelectEnvironment;
