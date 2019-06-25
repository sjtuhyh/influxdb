import * as React from 'react'
import {render, fireEvent} from 'react-testing-library'
import ThresholdsSettings from 'src/shared/components/ThresholdsSettings'
import {BASE_THRESHOLD_ID} from 'src/shared/constants/thresholds'

describe('ThresholdSettings', () => {
  const getErrorMessage = (container, thresholdID) => {
    const node = container.querySelector(
      `.threshold-setting[data-test-id='${thresholdID}'] .threshold-setting--error`
    )

    return node ? node.textContent.trim() : null
  }

  const getInput = (container, thresholdID) =>
    container.querySelector(
      `.threshold-setting[data-test-id='${thresholdID}'] input`
    )

  test('making then correcting an error', () => {
    const thresholds = [
      {
        id: BASE_THRESHOLD_ID,
        type: 'threshold',
        name: 'thunder',
        hex: '',
        value: null,
      },
      {id: '0', type: 'threshold', name: 'fire', hex: '', value: 30},
    ]

    const {container} = render(
      <ThresholdsSettings thresholds={thresholds} onSetThresholds={jest.fn()} />
    )

    // Enter an invalid value in the input
    fireEvent.change(getInput(container, '0'), {
      target: {value: 'baloney'},
    })

    // Blur the input
    fireEvent.blur(getInput(container, '0'))

    // Expect an error message to exist
    expect(getErrorMessage(container, '0')).toEqual(
      'Please enter a valid number'
    )

    // Enter a valid value in the input
    fireEvent.change(getInput(container, '0'), {
      target: {value: '9000'},
    })

    // Blur the input
    fireEvent.blur(getInput(container, '0'))

    // Expect there to be no error
    expect(getErrorMessage(container, '0')).toBeNull()
  })

  test('entering value less than min threshold shows error', () => {
    const thresholds = [
      {id: '0', type: 'min', name: 'thunder', hex: '', value: 20},
      {id: '1', type: 'threshold', name: 'fire', hex: '', value: 30},
      {id: '2', type: 'max', name: 'ruby', hex: '', value: 60},
    ]

    const {container} = render(
      <ThresholdsSettings thresholds={thresholds} onSetThresholds={jest.fn()} />
    )

    // Enter a value in the input
    fireEvent.change(getInput(container, '1'), {
      target: {value: '10'},
    })

    // Blur the input
    fireEvent.blur(getInput(container, '1'))

    // Expect an error message to exist
    expect(getErrorMessage(container, '1')).toEqual(
      'Please enter a value greater than the minimum threshold'
    )
  })

  test('entering value greater than max threshold shows error', () => {
    const thresholds = [
      {id: '0', type: 'min', name: 'thunder', hex: '', value: 20},
      {id: '1', type: 'threshold', name: 'fire', hex: '', value: 30},
      {id: '2', type: 'max', name: 'ruby', hex: '', value: 60},
    ]

    const {container} = render(
      <ThresholdsSettings thresholds={thresholds} onSetThresholds={jest.fn()} />
    )

    // Enter a value in the input
    fireEvent.change(getInput(container, '1'), {
      target: {value: '80'},
    })

    // Blur the input
    fireEvent.blur(getInput(container, '1'))

    // Expect an error message to be called
    expect(getErrorMessage(container, '1')).toEqual(
      'Please enter a value less than the maximum threshold'
    )
  })
})
