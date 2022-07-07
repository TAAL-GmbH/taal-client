export const colDefs = [
  {
    id: 'created_at',
    name: 'Created at (UTC)',
    type: 'dateStr',
  },
  {
    id: 'id',
    name: 'ID',
    type: 'transactionhash',
  },
  {
    id: 'api_key',
    name: 'API key used',
    type: 'apiKey',
  },
  {
    id: 'data_bytes',
    name: 'Data size',
    type: 'dataSize',
  },
  {
    id: 'filename',
    name: 'Filename',
    type: 'string',
  },
]

export const rangeItems = [
  { label: '24H', value: '24' },
  { label: '1W', value: '168' },
  { label: '1M', value: '720' },
  { label: 'All', value: '' },
]
