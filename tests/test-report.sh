#!/bin/bash

set -e

curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-001",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'

curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-001",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'

  curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-001",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'

  curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-001",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'

  curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-003",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'

  curl -X POST \
  http://localhost:8080/api/v1/reports/exam \
  -H "Content-Type: application/json" \
  --output report.pdf \
  --data '{
    "student": {
      "id": "STU-002",
      "name": "Arun Kumar",
      "rollNumber": "10A01",
      "class": "10-A"
    },
    "exam": {
      "id": "EXAM-2026-001",
      "name": "Final Examination",
      "year": 2026,
      "subjects": [
        {
          "name": "Mathematics",
          "marks": 89,
          "maxMarks": 100
        },
        {
          "name": "Science",
          "marks": 94,
          "maxMarks": 100
        },
        {
          "name": "English",
          "marks": 82,
          "maxMarks": 100
        }
      ]
    }
  }'
echo ""
echo "Report generated: report.pdf"