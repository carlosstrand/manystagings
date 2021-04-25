import * as React from 'react';
import { useHistory, useParams } from 'react-router-dom';
import CreateEditEnvVarModal from '../CreateEditEnvVarModal';
import { DataGrid, GridColDef } from '@material-ui/data-grid';
import ApplicationEnvVar from '../../types/application_env_var';

const columns: GridColDef[] = [
  { field: 'key', headerName: 'Key', flex: 0.1 },
  { field: 'value', headerName: 'Value', flex: 0.1 },
];

const EnvVars = ({ envVars }: { envVars: ApplicationEnvVar[] }) => {
  const [isEditingId, setIsEditingId] = React.useState(null);
  const history = useHistory();
  const { envId } = useParams();
  if (!envVars) {
    return null;
  }
  return (
    <>
      <CreateEditEnvVarModal id={isEditingId} open={!!isEditingId} onClose={() => setIsEditingId(null)} />
      <DataGrid
        rows={envVars}
        columns={columns}
        pageSize={10}
        autoHeight
        disableColumnSelector
        disableSelectionOnClick
        onRowClick={(row) => {
          // history.push(`/environments/${envId}/applications/${row.id}`);
          setIsEditingId(row.id);
        }}
      />
    </>
  );
}

export default EnvVars;
