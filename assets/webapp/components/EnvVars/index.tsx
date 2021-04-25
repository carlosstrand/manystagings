import * as React from 'react';
import { useHistory, useParams } from 'react-router-dom';
import CreateEditEnvVarModal from '../CreateEditEnvVarModal';
import { DataGrid, GridColDef } from '@material-ui/data-grid';
import ApplicationEnvVar from '../../types/application_env_var';
import useEnvVars from '../../hooks/useEnvVars';
import styled from 'styled-components';
import { StyledButton } from '../../ui/Button/style';
import Button from '../../ui/Button';
import Add from '@material-ui/icons/Add';
import IconButton from '@material-ui/core/IconButton';
import { Delete, Edit } from '@material-ui/icons';
import useDeleteEnvVar from '../../hooks/useDeleteEnvVar';


const Header = styled.div`
  margin-bottom: 16px;
  clear: both;
  h2 {
    display: inline-block;
  }
  ${StyledButton} {
    float: right;
  }
`;

const EnvVars = () => {
  const [isEditingId, setIsEditingId] = React.useState(null);
  const [addModalOpen, setAddModalOpen] = React.useState(false);
  const history = useHistory();
  const { appId } = useParams();
  const envQuery = useEnvVars(appId);
  const deleteQuery = useDeleteEnvVar({});
  const envVars = envQuery.data?.data;
  if (!envVars) {
    return null;
  }
  const columns: GridColDef[] = [
    { field: 'key', headerName: 'Key', flex: 0.1 },
    { field: 'value', headerName: 'Value', flex: 0.1 },
    {
      field: 'options',
      width: 200,
      renderCell: (params) => {
        return (
          <>
            <IconButton onClick={(e) => {
              e.preventDefault();
              setIsEditingId(params.id);
            }}>
              <Edit />
            </IconButton>
            <IconButton onClick={(e) => {
              e.preventDefault();
              const confirmed = confirm(`Are you sure you want to delete the key ${params.getValue("key")}?`);
              if (confirmed) {
                deleteQuery.mutateAsync({
                  id: params.getValue("id").toString(),
                }).then(() => {
                  envQuery.refetch();
                })
              }
            }}>
              <Delete />
            </IconButton>
          </>
        )
      }
    }
  ];
  return (
    <>
      {!!isEditingId && (
        <CreateEditEnvVarModal
          id={isEditingId}
          open={!!isEditingId}
          onClose={() => setIsEditingId(null)}
          refetch={envQuery.refetch}
        />
      )}
      {addModalOpen && (
        <CreateEditEnvVarModal
          id={null}
          open={addModalOpen}
          onClose={() => setAddModalOpen(null)}
          refetch={envQuery.refetch}
        />
      )}
      <Header>
        <h2>Environment Variables</h2>
        <Button onClick={() => setAddModalOpen(true)}>
          <Add /> Add
        </Button>
      </Header>
      <DataGrid
        rows={envVars}
        columns={columns}
        pageSize={30}
        autoHeight
        disableColumnSelector
        disableSelectionOnClick
      />
    </>
  );
}

export default EnvVars;
