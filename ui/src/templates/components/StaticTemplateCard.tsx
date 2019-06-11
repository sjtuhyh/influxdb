// Libraries
import React, {PureComponent, MouseEvent} from 'react'
import _ from 'lodash'
import {connect} from 'react-redux'
import {withRouter, WithRouterProps} from 'react-router'
import {
  Button,
  ComponentSize,
  ComponentSpacer,
  FlexDirection,
  JustifyContent,
} from '@influxdata/clockface'

// Components
import {ResourceList} from 'src/clockface'

// Actions
import {createResourceFromStaticTemplate} from 'src/templates/actions'

// Selectors
import {viewableLabels} from 'src/labels/selectors'

// Types
import {TemplateSummary, ILabel} from '@influxdata/influx'
import {ComponentColor} from '@influxdata/clockface'
import {AppState, Organization} from 'src/types'

// Constants
interface OwnProps {
  template: TemplateSummary
  name: string
  onFilterChange: (searchTerm: string) => void
}

interface DispatchProps {
  onCreateFromTemplate: typeof createResourceFromStaticTemplate
}

interface StateProps {
  labels: ILabel[]
  org: Organization
}

type Props = DispatchProps & OwnProps & StateProps

class StaticTemplateCard extends PureComponent<Props & WithRouterProps> {
  public render() {
    const {template} = this.props

    return (
      <ResourceList.Card
        testID="template-card"
        contextMenu={() => this.contextMenu}
        description={() => (
          <ResourceList.Description
            description={_.get(template, 'content.data.attributes.description')}
          />
        )}
        name={() => (
          <ResourceList.Name
            onClick={this.handleNameClick}
            name={template.meta.name}
            parentTestID="template-card--name"
            buttonTestID="template-card--name-button"
            inputTestID="template-card--input"
          />
        )}
      />
    )
  }

  private get contextMenu(): JSX.Element {
    return (
      <ComponentSpacer
        margin={ComponentSize.Medium}
        direction={FlexDirection.Row}
        justifyContent={JustifyContent.FlexEnd}
      >
        <Button
          text="Create"
          color={ComponentColor.Primary}
          size={ComponentSize.ExtraSmall}
          onClick={this.handleCreate}
        />
      </ComponentSpacer>
    )
  }

  private handleCreate = () => {
    const {onCreateFromTemplate, name} = this.props

    onCreateFromTemplate(name)
  }

  private handleNameClick = (e: MouseEvent<HTMLAnchorElement>) => {
    e.preventDefault()
    this.handleViewTemplate()
  }

  private handleViewTemplate = () => {
    const {router, org, name} = this.props

    router.push(`/orgs/${org.id}/templates/${name}/static/view`)
  }
}

const mstp = ({labels, orgs: {org}}: AppState): StateProps => {
  return {
    org,
    labels: viewableLabels(labels.list),
  }
}

const mdtp: DispatchProps = {
  onCreateFromTemplate: createResourceFromStaticTemplate,
}

export default connect<StateProps, DispatchProps, OwnProps>(
  mstp,
  mdtp
)(withRouter<Props>(StaticTemplateCard))
