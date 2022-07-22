const pageKey = 'page.history'

export const getColDefs = (t) => {
  return [
    {
      id: 'created_at',
      name: t(`${pageKey}.col-defs-label.created_at`),
      type: 'dateStr',
      props: {
        width: '20%',
      },
    },
    {
      id: 'id',
      name: t(`${pageKey}.col-defs-label.id`),
      type: 'transactionhash',
      props: {
        width: '14%',
      },
    },
    {
      id: 'api_key',
      name: t(`${pageKey}.col-defs-label.api_key`),
      type: 'apiKey',
      props: {
        width: '38%',
        style: {
          'grid-column-start': 1,
          'grid-column-end': -1,
        },
      },
    },
    {
      id: 'data_bytes',
      name: t(`${pageKey}.col-defs-label.data_bytes`),
      type: 'dataSize',
      props: {
        width: '10%',
      },
    },
    {
      id: 'filename',
      name: t(`${pageKey}.col-defs-label.filename`),
      type: 'filename',
      props: {
        width: '18%',
      },
    },
  ]
}

export const getRangeItems = (t) => {
  return [
    { label: t(`${pageKey}.range-label.24h`), value: '24' },
    { label: t(`${pageKey}.range-label.1w`), value: '168' },
    { label: t(`${pageKey}.range-label.1m`), value: '720' },
    { label: t(`${pageKey}.range-label.all`), value: '' },
  ]
}
