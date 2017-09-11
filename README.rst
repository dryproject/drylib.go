*************
DRYlib for Go
*************

.. image:: https://img.shields.io/badge/license-Public%20Domain-blue.svg
   :alt: Project license
   :target: https://unlicense.org/

.. image:: https://img.shields.io/badge/godoc-reference-blue.svg
   :alt: GoDoc reference
   :target: https://godoc.org/github.com/dryproject/drylib.go

.. image:: https://goreportcard.com/badge/github.com/dryproject/drylib.go
   :alt: Go Report Card score
   :target: https://goreportcard.com/report/github.com/dryproject/drylib.go

.. image:: https://img.shields.io/travis/dryproject/drylib.go/master.svg
   :alt: Travis CI build status
   :target: https://travis-ci.org/dryproject/drylib.go

|

http://drylib.org

Prerequisites
=============

* `Go <https://en.wikipedia.org/wiki/Go_(programming_language)>`__
  1.9+ (due to use of type aliases)

Features
========

Caveats
=======

Installation
============

::

   $ go get github.com/dryproject/drylib.go

Usage
=====

To import the library, do::

   import "github.com/dryproject/drylib.go"

Reference
=========

``core``
--------

================= ==============================================================
DRY Symbol        Go Symbol
================= ==============================================================
``core.bool``     ``dry.Bool`` (type alias for ``bool``)
``core.char``     ``dry.Char`` (type alias for ``rune``)
``core.complex``  ``dry.Complex`` (struct)
``core.float``    ``dry.Float`` (type alias for ``float64``)
``core.float32``  ``dry.Float32`` (type alias for ``float32``)
``core.float64``  ``dry.Float64`` (type alias for ``float64``)
``core.int``      ``dry.Int`` (type alias for ``int``)
``core.int8``     ``dry.Int8`` (type alias for ``int8``)
``core.int16``    ``dry.Int16`` (type alias for ``int16``)
``core.int32``    ``dry.Int32`` (type alias for ``int32``)
``core.int64``    ``dry.Int64`` (type alias for ``int64``)
``core.int128``   ``dry.Int128`` (type alias for ``Integer``)
``core.integer``  ``dry.Integer`` (struct)
``core.natural``  ``dry.Natural`` (type alias for ``Integer``)
``core.rational`` ``dry.Rational`` (struct)
``core.real``     ``dry.Real`` (struct)
``core.word``     ``dry.Word`` (type alias for ``uint``)
``core.word8``    ``dry.Word8`` (type alias for ``uint8``)
``core.word16``   ``dry.Word16`` (type alias for ``uint16``)
``core.word32``   ``dry.Word32`` (type alias for ``uint32``)
``core.word64``   ``dry.Word64`` (type alias for ``uint64``)
================= ==============================================================
