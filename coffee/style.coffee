# Public Domain (-) 2015 The Wikifactory.net Website Authors.
# See the Wikifactory.net UNLICENSE file for details.

module.exports = (api) ->

    api.add
        # Override default box model.
        '*':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'

        # Prevent iOS text size adjust after orientation change, without disabling user zoom.
        html:
            MsTextSizeAdjust: '100%'
            WebkitTextSizeAdjust: '100%'

        # Improve readability when focused and mouse hovered on links in all browsers.
        a:hover:
            outline: 0

        # Global styling.
        body:
            margin: 0
            padding: 0
            paddingTop: '60px'
            textRendering: 'optimizeLegibility'

        h1:
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: ''
            color: ''
            letterSpacing: ''

        h2:
            fontFamily: 'Montserrat'
            fontSize: '26px'
            fontWeight: 400
            lineHeight: 1.3
            color: '#2b2b2b'
            textAlign: 'center'
            marginTop: '50px'
            marginBottom: '50px'
            clear: 'both'

        h3:
            fontFamily:'Merriweather Sans'
            fontSize: '21px'
            fontWeight: 400
            lineHeight: 1.3
            marginBottom: 0

        p:
            fontFamily: 'Merriweather'
            fontSize:'14px'
            fontWeight: 300
            padding: 0
            margin: 0

        a:
            fontFamily: 'Merriweather'
            fontSize:'14px'
            fontWeight: 400
            color: '#cc0000'
            textDecoration: 'none'

        'a::selection':
           color: '#fff'

        '::selection':
            background: 'rgba(204,0,0,0.7)'

        tbody:
            # borderSpacing: '0px !important'
            # borderCollapse: 'collapse'
            width: '100%'

        tr:
            width: '100%'

        td:
            paddingTop: '0px'
            marginTop: '0px'

        footer:
            clear:'both'

        # HEADER
         header:
            width: '100%'
            height: '60px'
            backgroundColor: '#fff'
            top: 0
            zIndex: 100
            transition: 'top 0.2s ease-in-out'
            position: 'fixed'
            boxShadow: '0 1px 10px rgba(0, 0, 0, 0.1)'

        '.nav-up':
            top: '-60px'

        'header img':
            width: '42px'
            height: 'auto'
            marginTop: '8px'
            marginLeft: '20px'
            opacity: 0.8

        'header img:hover':
            cursor: 'pointer'
            opacity: 0.9

        # PARTICLES AND TAGLINE
        canvas:
            display: 'block'
            verticalAlign: 'bottom'

        '#particles-js':
            backgroundColor: '#0274B8' # Wikifactory Blue.
            color: '#ffffff'
            height: '500px'

        # MAIN WRAPPER
        '.wrapper':
            margin: '0 auto'
            width: '990px'
            clear: 'left'
