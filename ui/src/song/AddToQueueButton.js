import React from 'react'
import { Button, useDataProvider, useUnselectAll } from 'react-admin'
import { useDispatch } from 'react-redux'
import { addTrack } from '../player'

const AddToQueueButton = ({ selectedIds }) => {
  const dispatch = useDispatch()
  const dataProvider = useDataProvider()
  const unselectAll = useUnselectAll()
  const addToQueue = () => {
    selectedIds.forEach((id) => {
      dataProvider.getOne('song', { id }).then((response) => {
        dispatch(addTrack(response.data))
      })
    })
    unselectAll('song')
  }

  return <Button color="secondary" label="Add To Queue" onClick={addToQueue} />
}

export default AddToQueueButton