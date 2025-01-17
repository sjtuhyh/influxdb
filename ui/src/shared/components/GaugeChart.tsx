// Libraries
import React, {PureComponent} from 'react'
import _ from 'lodash'
import {AutoSizer} from 'react-virtualized'

// Components
import Gauge from 'src/shared/components/Gauge'

// Types
import {GaugeView} from 'src/types/dashboards'

import {ErrorHandling} from 'src/shared/decorators/errors'

interface Props {
  value: number
  properties: GaugeView
}

@ErrorHandling
class GaugeChart extends PureComponent<Props> {
  public render() {
    const {value} = this.props
    const {colors, prefix, suffix, decimalPlaces} = this.props.properties

    return (
      <AutoSizer>
        {({width, height}) => (
          <div className="gauge">
            <Gauge
              width={width}
              height={height}
              colors={colors}
              prefix={prefix}
              suffix={suffix}
              gaugePosition={value}
              decimalPlaces={decimalPlaces}
            />
          </div>
        )}
      </AutoSizer>
    )
  }
}

export default GaugeChart
