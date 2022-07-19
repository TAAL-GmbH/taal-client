export const colDefs = [
  {
    id: 'created_at',
    name: 'Created at (UTC)',
    type: 'dateStr',
    props: {
      width: '18%',
    },
  },
  {
    id: 'id',
    name: 'ID',
    type: 'transactionhash',
    props: {
      width: '17%',
    },
  },
  {
    id: 'api_key',
    name: 'API key used',
    type: 'apiKey',
    props: {
      width: '35%',
      style: {
        'grid-column-start': 1,
        'grid-column-end': -1,
      },
    },
  },
  {
    id: 'data_bytes',
    name: 'Data size',
    type: 'dataSize',
    props: {
      width: '10%',
    },
  },
  {
    id: 'filename',
    name: 'Filename',
    type: 'string',
    props: {
      width: '20%',
    },
  },
]

export const rangeItems = [
  { label: '24H', value: '24' },
  { label: '1W', value: '168' },
  { label: '1M', value: '720' },
  { label: 'All', value: '' },
]
