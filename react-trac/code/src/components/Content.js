import React from 'react'
import ContentHeader from './ContentHeader'
import ContentMain from './ContentMain'

export default function Content() {
    return (
        <div className="content-wrapper">
            <ContentHeader></ContentHeader>
            <ContentMain></ContentMain>
        </div>
    )
}
