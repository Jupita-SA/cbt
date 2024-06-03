document.addEventListener('DOMContentLoaded', () => {
  loadQuestion(0)
})

async function loadQuestion (index) {
  try {
    const response = await fetch(`/question?index=${index}`, {
      headers: { Accept: 'application/json' }
    })
    if (!response.ok) {
      throw new Error(`Failed to fetch question: ${response.statusText}`)
    }
    const data = await response.json()
    updateQuestionContent(data)
  } catch (error) {
    console.error('Error:', error)
    alert('Failed to load question. Please try again.')
  }
}

function updateQuestionContent (data) {
  const questionContainer = document.getElementById('question-container')
  const questionNumbersContainer = document.getElementById(
    'question-numbers-container'
  )

  let questionHTML = `
    <h1 class="card-title">Question ${data.index + 1} / ${data.total}</h1>
    <p class="card-text">${data.question.text}</p>
  `

  if (data.question.image) {
    questionHTML += `<img src="${data.question.image}" alt="Question Image" class="img-fluid mb-3" />`
  }

  questionHTML += `<form id="question-form">
    <input type="hidden" name="index" value="${data.index}" />
    <input type="hidden" name="navigation" id="navigation" value="" />
    <div class="form-group">`

  if (data.question.type === 'multiple_choice') {
    data.question.choices.forEach((choice, index) => {
      questionHTML += `
        <div class="form-check">
          <input class="form-check-input" type="radio" name="answer" value="${index}" id="choice${index}" />
          <label class="form-check-label" for="choice${index}">${choice}</label>
        </div>`
    })
  } else if (data.question.type === 'true_false') {
    questionHTML += `
      <div class="form-check">
        <input class="form-check-input" type="radio" name="answer" value="true" id="true" />
        <label class="form-check-label" for="true">True</label>
      </div>
      <div class="form-check">
        <input class="form-check-input" type="radio" name="answer" value="false" id="false" />
        <label class="form-check-label" for="false">False</label>
      </div>`
  } else if (data.question.type === 'short_answer') {
    questionHTML += `<input type="text" name="answer" class="form-control" placeholder="Your answer" />`
  }

  questionHTML += `</div>
    <div class="d-flex justify-content-between">`

  if (data.index > 0) {
    questionHTML += `<button type="button" class="btn btn-secondary" onclick="navigate('prev')">Prev</button>`
  }

  if (data.index < data.total - 1) {
    questionHTML += `<button type="button" class="btn btn-secondary" onclick="navigate('next')">Next</button>`
  } else {
    questionHTML += `<button type="button" class="btn btn-success" onclick="submitQuiz()">Submit</button>`
  }

  questionHTML += `</div></form>`
  questionContainer.innerHTML = questionHTML

  let questionNumbersHTML = ''
  data.questionNumbers.forEach((number, index) => {
    questionNumbersHTML += `
      <button type="button" class="btn btn-outline-primary m-1" style="width: 43px" onclick="loadQuestion(${index})">
        ${index + 1}
      </button>`
  })
  questionNumbersContainer.innerHTML = questionNumbersHTML

  maintainAnswerState(data.index, data.question.type)
}

function maintainAnswerState (currentQuestion, questionType) {
  const answers = JSON.parse(localStorage.getItem('answers')) || {}
  const questionIndex = `question${currentQuestion}`

  if (questionType === 'short_answer') {
    const input = document.querySelector('input[name="answer"]')
    if (answers[questionIndex]) {
      input.value = answers[questionIndex]
    }
    input.addEventListener('input', () => {
      answers[questionIndex] = input.value
      localStorage.setItem('answers', JSON.stringify(answers))
    })
  } else if (
    questionType === 'multiple_choice' ||
    questionType === 'true_false'
  ) {
    document.querySelectorAll('input[name="answer"]').forEach(input => {
      if (answers[questionIndex] === input.value) {
        input.checked = true
      }
      input.addEventListener('change', () => {
        answers[questionIndex] = input.value
        localStorage.setItem('answers', JSON.stringify(answers))
      })
    })
  }
}

async function navigate (action) {
  try {
    const form = document.getElementById('question-form')
    const formData = new FormData(form)
    formData.set('navigation', action)

    const response = await fetch('/question', {
      method: 'POST',
      body: formData
    })
    if (!response.ok) {
      throw new Error(`Failed to navigate: ${response.statusText}`)
    }
    const data = await response.json()
    loadQuestion(data.index)
  } catch (error) {
    console.error('Error:', error)
    alert('Failed to navigate. Please try again.')
  }
}

function submitQuiz () {
  window.location.href = '/result'
}
