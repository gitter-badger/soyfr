module.exports = function(grunt){

  'use strict';

  // Force use of Unix newlines
  grunt.util.linefeed = '\n';

  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');

  // Project configuration.
  grunt.initConfig({

    // Metadata.
    pkg: grunt.file.readJSON('package.json'),

    less: {
      compileCore: {
        options: {
          strictMath: true,
          noColor: true,
          outputSourceFiles: true
        },
        files: {
          'public/css/<%= pkg.name %>.css': 'public/less/bootstrap.less'
        }
      },
      minify: {
        options: {
          cleancss: true,
          report: 'min'
        },
        files: {
          'public/css/<%= pkg.name %>.min.css': 'public/css/<%= pkg.name %>.css'
        }
      }
    },

    copy: {
      fonts: {
        src: 'public/bower/bootstrap/fonts/*',
        dest: 'public/fonts',
        expand: true,
        flatten: true,
        filter: 'isFile'
      }
    },

    watch: {
      less: {
        files: '**/*.less',
        tasks: ['less']
      }
    }
  });

  grunt.registerTask('copyfonts',  ['copy:fonts']);
  grunt.registerTask('default', ['less', 'copy']);
};
