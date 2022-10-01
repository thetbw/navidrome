import React from 'react'
import { useMediaQuery } from '@material-ui/core'
import { useTranslate } from 'react-admin'

export const Title = ({ subTitle, args }) => {
  const translate = useTranslate()
  const isDesktop = useMediaQuery((theme) => theme.breakpoints.up('md'))
  const text = translate(subTitle, { ...args, _: subTitle })

  if (isDesktop) {
    return <span>白小耀 {text ? ` - ${text}` : ''}</span>
  }
  return <span>{text ? text : '白小耀'}</span>
}
